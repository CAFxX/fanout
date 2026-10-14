package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 int32) uint {
    fmt.Println(p0, p0)
    if (!(0.1 >= 0.1)) {
        v4 := GenS1{F0: 2}
        var v5 uint8 = 1
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS2{F0: "\n\t"}
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 857
}

func main() {
    var v1 int64 = 0
    var v2 float32 = float32(((1.5 * (0.0 / (0.1 / 1.5))) - -2.25))
    var v3 int16 = -524
    switch "\n\t" {
    case "hello":
        fmt.Println(v2, v1, v3)
    case "\n\t":
        var v8 uint32 = 4294967295
        _ = v8
    case "":
        for fi9 := 0; fi9 < 3; fi9++ {
            switch 32768 {
            case -264:
                if false {
                    var v10 uint8 = 2
                    v11 := []int{(fi9 / ((fi9) | 1)), 1, (fi9 & fi9)}
                    _ = v10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 9; fi12++ {
                    fmt.Println(fi12)
                    switch (^(fi12)) {
                    case 32768:
                        fmt.Println(v2, fi12, v3)
                    default:
                        var v13 uint = genFn0(842, 1)
                        _ = v13
                    }
                    _ = fi12
                }
            }
            _ = fi9
        }
    }
    var v14 uint32 = 730
    switch 42 {
    case -2147483648:
        if false {
            var v15 int64 = v1
            var v16 uint16 = 0
            _ = v15
            _ = v16
        } else {
            var v17 uint = 65535
            fmt.Println(v3, v17)
            v18 := GenS1{F0: 128}
            _ = v17
            _ = v18
        }
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        var v19 uint8 = 171
        _ = v19
    }
    v3 = v3
    if (0 < 2147483647) {
        var v20 uint16 = 42
        v21 := GenS1{F0: 42}
        v22 := []int{-2147483648, 9223372036854775807}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 42}
    var v24 float32 = float32((3.14159 - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
