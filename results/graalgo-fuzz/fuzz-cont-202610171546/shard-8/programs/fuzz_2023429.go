package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0() uint16 {
    v7 := []int{-2147483648, 32768, 36, -1}
    var v8 int16 = 0
    var v9 float32 = float32(-2.25)
    fmt.Println(v8, v9)
    v9 += ((3.14159 * (3.14159 - (float32((1.5 + 3.14159)) + 0.0))) * (((float32(0.1) * v9) + v9) / v9))
    _ = v7
    _ = v8
    _ = v9
    return 362
}
func genFn1(p0 bool) int32 {
    var v10 uint32 = 0
    for fi11 := 0; fi11 < 1; fi11++ {
        switch (("\n\t" + "a") + ("\n\t" + "")) {
        case "":
            switch ((("\n\t" + "\n\t") + ((("foo bar" + "a") + ((("\n\t" + "") + "hello") + "")) + "a")) + ("" + ("hello" + "hello"))) {
            case "foo bar":
                v12 := []int{(fi11 ^ (-(fi11))), 4294967295}
                _ = v12
            default:
                var v13 uint16 = genFn0()
                _ = v13
            }
        default:
            var v14 string = ""
            _ = v14
        }
        _ = fi11
    }
    var v15 float64 = (0.1 / (1e10 + float64(3.14159)))
    fmt.Println(p0, v15)
    _ = p0
    _ = v10
    _ = v15
    return -2147483648
}

func main() {
    var v1 string = ""
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 int8 = -59
    v5 := make(map[string]string)
    var v6 uint32 = 1
    fmt.Println(v6, v2)
    var v16 float64 = (1.5 - -2.25)
    switch 127 {
    case -2147483648:
        v17 := []int{1}
        _ = v17
    case 0:
        v18 := GenS1{F0: genFn0()}
        _ = v18
    case 32767:
        var v19 uint16 = 0
        _ = v19
    default:
        v20 := []int{0, 2147483647, -375}
        _ = v20
    }
    v21 := v3[42]
    var v22 uint16 = 0
    v16 = 3.14159
    var v23 int8 = v4
    v24 := GenS1{F0: (v22 << ((0) & 63))}
    v21 = ((v1 + v21) + (("\n\t" + "a") + "\n\t"))
    v25 := GenS2{F0: 32767, F1: 1}
    v23 = v4
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS1{F0: (v22 + (v22 * (v22 << ((0) & 63))))}
        var v28 float32 = 1.5
        _ = fi26
        _ = v27
        _ = v28
    }
    v1 = v21
    if v2 {
        v5[v21] = "a"
    }
    if false {
        v29 := GenS3{F0: 127}
        var v30 int = -950
        if (true && (!(!v2))) {
            v6 += (v6 - 560)
            var v31 int8 = (v23 ^ -1)
            _ = v31
        }
        _ = v29
        _ = v30
    } else {
        v16 = v16
        v4 = 35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
