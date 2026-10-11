package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 bool) string {
    p0 = (255 >= 0)
    if true {
        v7 := GenS1{F0: 32767, F1: "", F2: true}
        var v8 uint = 303
        v9 := []int{32768, -9223372036854775808}
        _ = v7
        _ = v8
        _ = v9
    } else {
        switch "hello" {
        case "\n\t":
            var v10 int16 = -1
            _ = v10
        case "a":
            p0 = ((32767 > -666) || p0)
        case "foo bar":
            var v11 uint64 = 1
            _ = v11
        }
    }
    fmt.Println(p0)
    var v12 int8 = 0
    _ = p0
    _ = v12
    return "a"
}
func genFn1(p0 uint, p1 int64) uint8 {
    fmt.Println(p0, p0)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    p1 -= (p1 >> (((p0 | (^(p0)))) & 63))
    switch 9223372036854775807 {
    case 356:
        if (1 <= 42) {
            fmt.Println(p1)
            if true {
                var v15 uint32 = uint32(p0)
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 float32 = (((((0.1 - (3.14159 - 0.0)) * -2.25) - 3.14159) + 1.5) / 1.5)
                    switch "hello" {
                    case "hello":
                        var v18 uint64 = 0
                        _ = v18
                    }
                    if ((false && (false && (false && false))) || false) {
                        var v19 uint8 = 2
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                if ((("\n\t" == "") || (-914 != 32767)) && (-32768 == -1)) {
                    var v20 uint64 = 2
                    v15 = uint32(v20)
                    _ = v20
                }
                _ = v15
            } else {
                v21 := []int{78}
                var v22 string = ("foo bar" + "foo bar")
                _ = v21
                _ = v22
            }
            var v23 int8 = -41
            _ = v23
        } else {
            p1 -= (p1 ^ p1)
            v24 := GenS1{F0: 1, F1: "a", F2: (!(2147483648 == 428))}
            for fi25 := 0; fi25 < 12; fi25++ {
                p1 -= p1
                v26 := []int{}
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    }
    _ = p0
    _ = p1
    return 55
}

func main() {
    var v1 float64 = (3.14159 - 3.14159)
    v2 := make(map[string]bool)
    var v3 string = ("a" + "a")
    var v4 int64 = -9223372036854775808
    var v5 uint = 88
    v6 := make(map[string]string)
    var v27 uint8 = 157
    fmt.Println(v4, v3, v27)
    v4 = 646
    v1 -= v1
    fmt.Println(v4, v4)
    v27 += (v27 & v27)
    switch (v3 + v3) {
    case "hello":
        switch -2147483648 {
        case -1:
            v3 = v3
        }
    }
    var v28 int64 = -1
    fmt.Println(v28, v27, v27)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
