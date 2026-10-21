package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() string {
    if false {
        if false {
            var v5 int8 = 0
            v6 := []int{9223372036854775807, 613, -2147483648, 32768}
            _ = v5
            _ = v6
        }
    } else {
        switch -887 {
        case 4294967295:
            v7 := []int{127, -173, -9223372036854775808}
            _ = v7
        case 127:
            switch ("\n\t" + ("hello" + "foo bar")) {
            case "\n\t":
                v8 := []int{609, 1, 4294967295, 2147483647}
                _ = v8
            case "a":
                var v9 uint64 = 4294967295
                _ = v9
            case "foo bar":
                var v10 int32 = 0
                _ = v10
            }
        case -2147483648:
            var v11 int64 = -1
            _ = v11
        }
        var v12 int8 = 100
        _ = v12
    }
    var v13 int64 = -541
    v14 := GenS1{F0: (!true)}
    for fi15 := 0; fi15 < 5; fi15++ {
        v13 = 9223372036854775807
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int32 = 0
            _ = fi16
            _ = v17
        }
        var v18 float64 = 0.0
        _ = fi15
        _ = v18
    }
    fmt.Println(v13)
    _ = v13
    _ = v14
    return ("hello" + ("hello" + ""))
}
func genFn1(p0 int32, p1 string, p2 uint32) int32 {
    p0 = p0
    p2 = (^(p2))
    v19 := []int{}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p0
}

func main() {
    var v1 int64 = 345
    var v2 int32 = 1
    v3 := []int{127}
    var v4 uint16 = 570
    v1 = (v1 + (v1 * v1))
    v4 = (v4 | 1)
    v4 -= 102
    switch ("a" + "hello") {
    case "\n\t":
        if (true || true) {
            for fi20 := 0; fi20 < 12; fi20++ {
                v2 -= (v2 & v2)
                v4 = 0
                v2 = v2
                _ = fi20
            }
        } else {
            v21 := GenS1{F0: false}
            v1 -= (v1 | v1)
            v22 := GenS1{F0: true}
            _ = v21
            _ = v22
        }
    case "a":
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS1{F0: ((1 == -128) || false)}
            _ = fi23
            _ = v24
        }
    }
    v3 = append(v3, 9223372036854775807)
    fmt.Println(v4, v2, v4)
    if len(v3) > 0 {
        v3[0] = 969
    }
    var v25 bool = (true && false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
