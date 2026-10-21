package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0() bool {
    v6 := []int{487, 42}
    if (false || false) {
        var v7 int16 = 42
        _ = v7
    } else {
        v8 := GenS1{F0: 32768, F1: 980, F2: 1}
        v9 := GenS1{F0: 0, F1: 255, F2: 1000}
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        if (!(2 == 290)) {
            var v10 string = (("\n\t" + ("foo bar" + "foo bar")) + "hello")
            var v11 int64 = -1
            _ = v10
            _ = v11
        }
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v12 bool = ("hello" <= "hello")
        _ = v12
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    var v13 uint8 = 0
    _ = v13
    _ = v6
    return ((float64((1e10 / 0.1)) + 0.1) < (((0.1 * 0.1) + 0.0) / (1e10 + (((1.5 - 1e10) - -2.25) + 0.0))))
}
func genFn1(p0 int16, p1 uint) uint64 {
    v14 := []int{32768, 212, 32768}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v14
    return 42
}

func main() {
    var v1 int16 = -865
    var v2 uint32 = 2147483648
    var v3 float64 = 1.5
    var v4 int = 0
    var v5 uint16 = 2
    if ((v4 | 0) > 32767) {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := []int{(v4 & 516), (fi15 - (fi15 * 32767)), v4}
            if (100 >= -43) {
                fmt.Println(fi15)
            } else {
                if true {
                    var v17 bool = ((0.1 + 1.5) == 1e10)
                    var v18 bool = ((!false) && (!v17))
                    _ = v17
                    _ = v18
                }
                var v19 float32 = ((3.14159 * ((3.14159 + (1.5 - float32(3.14159))) * 1.5)) / 1e10)
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        v20 := GenS2{F0: (v5 * (^(v5))), F1: 42}
        _ = v20
    }
    var v21 string = "a"
    var v22 int32 = 42
    v23 := []int{127, 32768, (^(v4)), (v4 << ((0) & 63))}
    v24 := GenS1{F0: v4, F1: 1, F2: v1}
    v21 = ("\n\t" + (v21 + (("\n\t" + "\n\t") + (v21 + v21))))
    var v25 float32 = ((float32((float64((-2.25 - (((0.1 + (1e10 / 0.1)) + -2.25) + -2.25))) * (v3 - (0.1 + (1.5 - v3))))) + 3.14159) - 0.1)
    var v26 bool = false
    if (0 >= -1) {
        v3 = (v3 + 0.1)
    }
    if v26 {
        fmt.Println(v21)
        v27 := GenS3{F0: ((("a" + ("foo bar" + (v21 + ("\n\t" + "")))) + "") + "")}
        var v28 uint64 = 18446744073709551615
        _ = v27
        _ = v28
    } else {
        v23 = append(v23, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
