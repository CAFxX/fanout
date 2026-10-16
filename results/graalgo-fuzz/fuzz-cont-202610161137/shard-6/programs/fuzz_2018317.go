package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() bool {
    var v5 int = -612
    v5 -= -9223372036854775808
    v6 := []int{675, (^(v5))}
    if false {
        v5 += 2147483647
        if len(v6) > 0 {
            v6[0] = v5
        }
        v5 -= v5
    } else {
        var v7 bool = (311 <= 919)
        var v8 uint = 0
        _ = v7
        _ = v8
    }
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := []int{}
    v2 := []int{32767, 303}
    var v3 string = ""
    var v4 int16 = 1000
    v9 := GenS1{F0: v3}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int = fi10
        fmt.Println(v4, v11, v11)
        var v12 bool = true
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = ("a" <= v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    v14 := GenS1{F0: "foo bar"}
    v4 = (v4 / ((v4) | 1))
    fmt.Println(v4, v3)
    if (v13 || (2 >= 633)) {
        v15 := GenS1{F0: (v3 + v3)}
        switch v3 {
        case "hello":
            _ = len(v2)
        case "a":
            v4 = (v4 + v4)
        case "":
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 int = fi16
                v18 := GenS1{F0: (v3 + (v3 + v3))}
                v17 = (-(v17))
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        if (301 < 47) {
            v3 = "\n\t"
            var v19 uint16 = 42
            if v13 {
                if (65535 == 2) {
                    var v20 int32 = -2147483648
                    _ = v20
                }
                var v21 uint32 = 42
                var v22 uint8 = 2
                _ = v21
                _ = v22
            }
            _ = v19
        }
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
