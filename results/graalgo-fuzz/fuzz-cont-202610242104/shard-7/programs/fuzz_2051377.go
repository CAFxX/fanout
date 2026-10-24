package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    switch 4294967295 {
    case 32767:
        var v6 float64 = (0.1 / 1e10)
        _ = v6
    case -634:
        var v7 float32 = 0.1
        _ = v7
    case 128:
        if true {
            v8 := GenS1{F0: -37}
            var v9 uint = 0
            v10 := []int{-664}
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 uint64 = 765
            fmt.Println(v11)
            _ = v11
        }
    }
    var v12 bool = (42 == 2)
    fmt.Println(v12)
    var v13 float32 = (-2.25 / 0.1)
    fmt.Println(v12)
    _ = v12
    _ = v13
    return 255
}

func main() {
    var v1 float32 = -2.25
    var v2 bool = (42 >= 1)
    var v3 string = ("" + (("foo bar" + "\n\t") + "foo bar"))
    var v4 float64 = 0.0
    var v5 uint16 = 32767
    if (v2 || v2) {
        fmt.Println(v2)
        if (v2 && v2) {
            var v14 bool = ((301 > 0) && (-2147483648 <= 1))
            _ = v14
        } else {
            switch v3 {
            case "foo bar":
                fmt.Println(v3, v2)
            }
        }
    }
    if true {
        var v15 float64 = (v4 * v4)
        fmt.Println(v3)
        _ = v15
    }
    v3 = v3
    var v16 uint8 = 128
    v2 = true
    v17 := GenS1{F0: 127}
    var v18 float64 = 1.5
    if ((!v2) || (!(1 < -128))) {
        var v19 string = (("" + v3) + "")
        if v2 {
            if (!(285 != 2)) {
                var v20 bool = (3.14159 < (((((0.1 * v18) * v4) - (v4 + -2.25)) + (((v18 - float64(((float32(((((float64(v1) + v4) / 1.0) - -2.25) * (((v18 + ((v18 + (0.0 - v4)) * -2.25)) + (((-2.25 / v4) - v18) + (v18 - v18))) / (3.14159 + v4)))) / v1) + (((3.14159 + ((float32(v4) - 3.14159) * (-2.25 - v1))) * v1) + -2.25)))) / 1e10) / v4)) * v4))
                _ = v20
            }
            var v21 bool = v2
            _ = v21
        }
        _ = v19
    } else {
        v2 = (v2 || false)
    }
    if true {
        v22 := []int{-2147483648, -1, 697}
        var v23 uint64 = 18446744073709551615
        _ = v22
        _ = v23
    } else {
        v24 := []int{}
        if v2 {
            fmt.Println(v2)
            v25 := GenS1{F0: -128}
            _ = v25
        } else {
            if (!true) {
                v26 := GenS1{F0: -1}
                var v27 string = v3
                _ = v26
                _ = v27
            }
        }
        v24 = append(v24, 127)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
