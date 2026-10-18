package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0(p0 bool, p1 int64, p2 int8) int8 {
    v4 := []int{1, 2147483647}
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 string = "foo bar"
    var v3 uint = 42
    if (("" + v2) != (("hello" + v2) + v2)) {
        v3 = v3
    }
    v1 -= 1000
    var v5 bool = ((false || (true || (!true))) && true)
    v5 = false
    v5 = (287 == -9223372036854775808)
    v6 := GenS1{F0: 2147483647}
    v5 = (255 >= 42)
    switch (v2 + v2) {
    case "hello":
        for fi7 := 0; fi7 < 10; fi7++ {
            v3 = (v3 % ((65535) | 1))
            v8 := []int{}
            _ = fi7
            _ = v8
        }
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2, v2)
        case "hello":
            v9 := GenS1{F0: 30}
            _ = v9
        }
    }
    v1 -= v1
    v2 = v2
    fmt.Println(v1, v5, v3)
    var v10 int16 = v1
    if ((100 > -128) && (v10 != v10)) {
        v11 := GenS2{F0: 1}
        _ = v11
    } else {
        v10 += 765
        v12 := []int{4294967295}
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v10, v10)
            var v14 uint16 = 539
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v10 = int16(v3)
    var v15 uint16 = 249
    for fi16 := 0; fi16 < 19; fi16++ {
        v5 = (127 > 0)
        for fi17 := 0; fi17 < 3; fi17++ {
            v15 -= v15
            _ = fi17
        }
        switch v2 {
        case "\n\t":
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint = 18446744073709551615
                _ = fi18
                _ = v19
            }
        case "a":
            v20 := []int{32768}
            _ = v20
        case "":
            var v21 int = (fi16 & fi16)
            _ = v21
        }
        _ = fi16
    }
    v1 -= -32768
    v22 := GenS3{F0: v3, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v5
    _ = v6
}
