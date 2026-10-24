package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) int64 {
    v5 := GenS1{F0: 2147483647}
    v6 := GenS1{F0: -1}
    v7 := []int{-1, -241, -650, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() int16 {
    var v8 int16 = 1000
    var v9 int64 = -1
    if (!((("a" + "foo bar") + "\n\t") > "foo bar")) {
        switch ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) {
        case "\n\t":
            var v10 uint64 = 0
            _ = v10
        case "":
            v11 := GenS1{F0: -1}
            _ = v11
        case "foo bar":
            v12 := []int{-443, 32767, 4294967295, -208}
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 42}
        _ = v13
    }
    _ = v8
    _ = v9
    return 1000
}

func main() {
    var v1 uint = 680
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 int32 = int32(v1)
    fmt.Println(v2, v1)
    v14 := v3["\n\t"]
    v3[("\n\t" + ("\n\t" + "hello"))] = (3.14159 == float64(float32(-2.25)))
    var v15 uint64 = 1
    v16 := []int{-2147483648, 600, 127, -112}
    v16 = append(v16, -9223372036854775808)
    v17 := v3[(((("" + "\n\t") + "a") + "\n\t") + ("hello" + "\n\t"))]
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v18 uint = 65535
        _ = v18
    case "":
        v2 = false
    case "a":
        switch "\n\t" {
        case "":
            v17 = false
        case "hello":
            if v17 {
                v14 = v17
                var v19 int8 = 127
                switch 4294967295 {
                case 42:
                    v20 := GenS1{F0: (^(v4))}
                    _ = v20
                default:
                    v15 += (v15 + 42)
                }
                _ = v19
            } else {
                _ = len(v3)
            }
        case "a":
            v2 = true
        default:
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(v2)
                _ = fi21
            }
        }
    default:
        var v22 int16 = -1
        _ = v22
    }
    if (879 <= 604) {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v4, v15)
            case "\n\t":
                v2 = v14
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                delete(v3, "\n\t")
                _ = fi24
            }
            _ = fi23
        }
    } else {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    v14 = (127 == 127)
    var v25 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
