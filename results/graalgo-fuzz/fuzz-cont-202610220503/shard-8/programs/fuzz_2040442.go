package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 int32, p1 uint8) uint8 {
    fmt.Println(p1, p1)
    if true {
        var v4 uint8 = (^(p1))
        fmt.Println(v4, p0, p0)
        _ = v4
    }
    _ = p0
    _ = p1
    return (p1 ^ (p1 >> ((0) & 63)))
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    var v3 int16 = 1
    v5 := GenS1{F0: v2}
    var v6 int16 = (v3 | v3)
    var v7 float64 = 1.5
    switch ((("a" + "a") + "hello") + "\n\t") {
    case "":
        switch ("\n\t" + ("\n\t" + "")) {
        case "foo bar":
            var v8 int = 119
            _ = v8
        default:
            switch -964 {
            case 127:
                var v9 int16 = (-(v3))
                _ = v9
            case 32767:
                v10 := []int{127, -1, int(v3), 127}
                _ = v10
            case -215:
                for fi11 := 0; fi11 < 13; fi11++ {
                    if false {
                        _ = len(v1)
                    }
                    _ = fi11
                }
            default:
                v12 := []int{-9223372036854775808, 128, -97}
                _ = v12
            }
        }
    default:
        v13 := []int{}
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := []int{fi14, fi14}
        var v16 string = "hello"
        if (true && true) {
            switch fi14 {
            case 4294967295:
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 uint32 = (v2 + v2)
                    _ = fi17
                    _ = v18
                }
            case 42:
                if (true && false) {
                    v19 := GenS1{F0: 42}
                    v16 = "a"
                    fmt.Println(v6, fi14)
                    _ = v19
                } else {
                    v6 = (v6 << ((637) & 63))
                    v20 := v1[v16]
                    var v21 string = v16
                    _ = v20
                    _ = v21
                }
            default:
                var v22 int8 = -92
                _ = v22
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v23 := []int{}
    if ((0.1 * -2.25) < 3.14159) {
        var v24 int8 = 42
        for fi25 := 0; fi25 < 11; fi25++ {
            v7 = (3.14159 * v7)
            for fi26 := 0; fi26 < 16; fi26++ {
                v23 = append(v23, fi26)
                _ = fi26
            }
            _ = fi25
        }
        _ = v24
    }
    v23 = append(v23, 2147483647)
    v7 -= v7
    v27 := GenS1{F0: 2147483647}
    if (false || true) {
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(fi28)
            _ = fi28
        }
        var v29 uint16 = 1
        _ = v29
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS1{F0: v2}
            v32 := GenS1{F0: v2}
            v3 -= (v6 + v3)
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            v34 := GenS2{F0: (fi33 - (fi33 * 128)), F1: false, F2: ((0.0 - 0.0) >= (1e10 / (0.1 * (-2.25 - 1.5))))}
            _ = fi33
            _ = v34
        }
        var v35 uint8 = 255
        _ = v35
    }
    var v36 uint64 = 18446744073709551615
    fmt.Println(v3)
    var v37 uint8 = 0
    v38 := GenS3{F0: 42, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v5
    _ = v6
    _ = v7
}
