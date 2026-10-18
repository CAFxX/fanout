package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8) string {
    fmt.Println(p0)
    p0 -= (^(p0))
    _ = p0
    return ("foo bar" + ("foo bar" + ""))
}
func genFn1(p0 int64) uint {
    var v7 float64 = (-2.25 * float64((0.1 * (float32(-2.25) / (float32(3.14159) / 1e10)))))
    v8 := []int{}
    _ = len(v8)
    if false {
        switch "\n\t" {
        case "\n\t":
            if (true || false) {
                _ = len(v8)
                var v9 uint64 = 18446744073709551615
                v7 = v7
                _ = v9
            }
        case "":
            v10 := GenS1{F0: 32767}
            _ = v10
        case "a":
            fmt.Println(p0)
        }
        p0 = p0
    } else {
        if (false && ((true && (938 > -2147483648)) && ("a" > "a"))) {
            v7 = 3.14159
        }
        fmt.Println(v7)
    }
    _ = p0
    _ = v7
    _ = v8
    return 428
}

func main() {
    var v1 uint64 = 805
    var v2 string = ("" + ("hello" + "\n\t"))
    v3 := make(map[int]string)
    var v4 uint16 = 10
    var v5 uint = 4294967295
    var v6 int = -9223372036854775808
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: -360}
        _ = fi11
        _ = v12
    }
    if false {
        if true {
            v2 = "hello"
            v4 += 2
            v6 = v6
        }
    }
    var v13 int32 = -1
    switch 9223372036854775807 {
    case -333:
        var v14 int8 = 42
        _ = v14
    }
    v1 -= 18446744073709551615
    var v15 uint32 = 790
    v3[v6] = v2
    for fi16 := 0; fi16 < 15; fi16++ {
        v4 = v4
        _ = fi16
    }
    v17 := v3[(v6 ^ (v6 % ((v6) | 1)))]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
