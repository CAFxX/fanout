package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
func genFn0(p0 int32, p1 string, p2 uint16) int16 {
    v4 := GenS1{F0: 795, F1: 100, F2: ((("hello" + p1) + p1) + (p1 + ""))}
    var v5 uint32 = 870
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1() string {
    v6 := []int{2147483647, 128, 899}
    v6 = append(v6, 4294967295)
    var v7 bool = true
    v8 := GenS1{F0: 2147483648, F1: 100, F2: ("hello" + (("" + "\n\t") + "foo bar"))}
    _ = v6
    _ = v7
    _ = v8
    return ("a" + "hello")
}

func main() {
    var v1 int16 = -32768
    v2 := []string{}
    var v3 uint = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= 4294967295
    v1 -= (^(v1))
    var v9 float32 = 1e10
    if true {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 2147483647, F1: 26, F2: ("foo bar" + "")}
            _ = fi10
            _ = v11
        }
    }
    var v12 int8 = 127
    if (0 == 541) {
        _ = len(v2)
    }
    fmt.Println(v3, v3, v12)
    if (925 != 2147483647) {
        var v13 uint8 = 127
        _ = v13
    } else {
        switch 266 {
        case -1:
            var v14 float32 = v9
            _ = v14
        default:
            var v15 int = 121
            _ = v15
        }
    }
    if (18446744073709551615 >= 42) {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int32 = 1
            fmt.Println(fi16, fi16)
            _ = fi16
            _ = v17
        }
        if (!(948 != 18446744073709551615)) {
            var v18 int = int(v12)
            _ = v18
        } else {
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int16 = 1000
                _ = fi19
                _ = v20
            }
            v1 += 32767
        }
    } else {
        v9 += (0.0 + 3.14159)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
