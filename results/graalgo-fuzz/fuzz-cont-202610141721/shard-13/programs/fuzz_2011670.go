package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 int16) uint16 {
    switch 418 {
    case 896:
        var v5 uint64 = p0
        _ = v5
    case 2147483647:
        switch -9223372036854775808 {
        case -2147483648:
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 int32 = 0
                _ = fi6
                _ = v7
            }
        case 61:
            var v8 int32 = 42
            _ = v8
        case 4294967295:
            var v9 float32 = (1.5 - ((0.1 / 1.5) * 1e10))
            _ = v9
        }
    case -103:
        v10 := GenS1{F0: 42}
        _ = v10
    }
    if false {
        p0 = 262
    }
    v11 := []int{0}
    p0 += (^(p0))
    var v12 int64 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 int32, p1 uint64) uint {
    var v13 float32 = -2.25
    var v14 uint32 = 0
    v13 -= float32(0.1)
    v15 := []int{-9223372036854775808, 107, 9223372036854775807}
    if true {
        fmt.Println(p1, p0)
        v15 = append(v15, 127)
    } else {
        switch -80 {
        case 2147483647:
            p1 = p1
        case 0:
            for fi16 := 0; fi16 < 10; fi16++ {
                switch (fi16 | int(p1)) {
                case 557:
                    var v17 float64 = ((((((3.14159 + (((((float64(v13) * float64((float32(1e10) * v13))) + -2.25) * (-2.25 - float64(((v13 - float32(3.14159)) / (float32(float64(v13)) * v13))))) / (1e10 + 0.1)) - 0.0)) * ((0.0 + 3.14159) / 0.1)) * 1.5) / (-2.25 - 3.14159)) + (0.1 - (float64(v13) - 0.1))) + 1e10)
                    _ = v17
                case 32767:
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        var v18 uint64 = (p1 ^ 983)
                        _ = v18
                    }
                case 434:
                    var v19 int32 = (^(p0))
                    _ = v19
                }
                var v20 float64 = 0.0
                _ = fi16
                _ = v20
            }
        default:
            var v21 int64 = 1
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 int8 = 127
    var v2 string = "hello"
    v3 := []string{"\n\t", v2, v2}
    var v4 string = ((("" + (v2 + (v2 + v2))) + "\n\t") + "hello")
    v3 = append(v3, (v2 + v2))
    _ = len(v3)
    var v22 uint32 = 2147483648
    var v23 bool = (0 != 4294967295)
    v4 = "foo bar"
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int8 = v1
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS1{F0: 0}
            var v28 uint = 0
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        v2 = v4
        _ = fi24
        _ = v25
    }
    v22 += (v22 * (v22 >> ((4294967295) & 63)))
    fmt.Println(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 = (9223372036854775807 > 9223372036854775807)
    var v29 float32 = -2.25
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 += float32(float64((v29 / 1.0)))
    if v23 {
        var v30 int64 = 42
        var v31 float64 = (1.5 / 3.14159)
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        if false {
            v23 = ((!(3.14159 < float64(((float32(0.0) / ((v29 + (v29 + v29)) + (float32(float64(v29)) * (((v29 / v29) * ((1e10 + v29) + 1e10)) * v29)))) + -2.25)))) || (127 >= 1))
            fmt.Println(v4)
        } else {
            v33 := GenS1{F0: -1}
            _ = v33
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint32 = v22
            _ = fi34
            _ = v35
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
