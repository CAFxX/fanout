package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 string) uint64 {
    var v7 bool = true
    var v8 float32 = (((3.14159 * 1.5) * float32((3.14159 + -2.25))) + (0.0 - ((1.5 * -2.25) * (3.14159 + 1e10))))
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint8 = 127
    var v3 uint8 = 81
    var v4 uint = 2
    var v5 int16 = 1000
    v6 := make(map[string]string)
    var v9 float32 = 1.5
    v10 := GenS1{F0: 32767}
    fmt.Println(v9)
    var v11 int = 9223372036854775807
    v12 := GenS1{F0: 1}
    var v13 int8 = 100
    fmt.Println(v1, v3)
    delete(v6, ("hello" + "a"))
    _ = len(v6)
    if ((("\n\t" == (("" + "a") + "")) && ((128 == 229) || true)) && false) {
        fmt.Println(v13)
        if true {
            var v14 int32 = -2147483648
            var v15 uint16 = 2
            _ = v14
            _ = v15
        } else {
            if true {
                var v16 uint = (v4 | v4)
                v9 = ((v9 + (((1e10 + float32(float64(3.14159))) - v9) + v9)) + (((3.14159 - (v9 - (float32((float64(v9) + (0.1 / 1e10))) * (v9 - v9)))) * (v9 * (((1.5 * v9) * (((v9 - v9) + (0.1 + 3.14159)) / (v9 + v9))) - v9))) * (((v9 / 1.5) - v9) * float32(1e10))))
                _ = v16
            }
            fmt.Println(v4, v4, v4)
            switch ("\n\t" + "") {
            case "a":
                var v17 int = v11
                _ = v17
            case "\n\t":
                for fi18 := 0; fi18 < 7; fi18++ {
                    v19 := v6["a"]
                    v19 = "foo bar"
                    _ = fi18
                    _ = v19
                }
            case "hello":
                v20 := GenS2{F0: -782, F1: v4}
                _ = v20
            }
        }
    }
    v21 := GenS3{F0: (float32(0.1) != v9)}
    var v22 uint32 = uint32(v4)
    var v23 int64 = 73
    var v24 uint = v4
    if v1 {
        var v25 uint16 = 65535
        _ = v25
    }
    v26 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
