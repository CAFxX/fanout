package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0() uint8 {
    var v6 int64 = 9223372036854775807
    var v7 uint = 1
    for fi8 := 0; fi8 < 14; fi8++ {
        if (173 <= 4294967295) {
            v6 = (v6 >> ((2) & 63))
        }
        var v9 int32 = 37
        var v10 uint64 = 902
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 int16) string {
    var v11 string = "hello"
    if true {
        switch 9223372036854775807 {
        case 32767:
            if (43 >= 42) {
                var v12 float32 = float32(-2.25)
                v12 = v12
                _ = v12
            }
        case -2147483648:
            var v13 string = "\n\t"
            _ = v13
        case 2147483647:
            switch v11 {
            case "a":
                var v14 int16 = p0
                _ = v14
            }
        }
    }
    var v15 float32 = (0.0 / 1.5)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            switch ("\n\t" + v11) {
            case "\n\t":
                var v18 int64 = 9223372036854775807
                _ = v18
            }
            var v19 float32 = (1e10 * v15)
            _ = fi17
            _ = v19
        }
        var v20 uint64 = 42
        v21 := []int{1, (fi16 | 128), fi16}
        _ = fi16
        _ = v20
        _ = v21
    }
    _ = p0
    _ = v11
    _ = v15
    return ((("hello" + (v11 + "\n\t")) + v11) + ((("hello" + ("hello" + "")) + (v11 + (((("\n\t" + v11) + ("foo bar" + v11)) + ((v11 + "foo bar") + (v11 + v11))) + (v11 + ((((v11 + ((v11 + ((((("\n\t" + "foo bar") + "foo bar") + v11) + "foo bar") + v11)) + v11)) + v11) + v11) + v11))))) + v11))
}

func main() {
    var v1 uint32 = 847
    var v2 int64 = -180
    var v3 string = "foo bar"
    var v4 int16 = -102
    var v5 bool = (!(0.1 == 0.1))
    switch -9223372036854775808 {
    case -9223372036854775808:
        v22 := []int{127}
        _ = v22
    case 32768:
        fmt.Println(v1, v1, v2)
    case -2147483648:
        var v23 int32 = 932
        _ = v23
    }
    v24 := GenS1{F0: 835}
    fmt.Println(v4, v2, v5)
    var v25 float64 = 1.5
    v26 := []int{14, -1, int(v4), int(v1)}
    var v27 int16 = (-(v4))
    v28 := GenS2{F0: 2}
    var v29 int8 = 100
    fmt.Println(v2, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
