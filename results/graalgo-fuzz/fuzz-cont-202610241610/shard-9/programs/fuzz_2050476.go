package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint16; F2 uint8 }

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 int64 = -49
    var v3 bool = (!(!true))
    v2 += (v2 / (((v2 / ((v2) | 1))) | 1))
    v4 := GenS1{F0: 4294967295, F1: 0}
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v3 = false
            v2 = 1
            var v7 int = (fi6 - fi5)
            _ = fi6
            _ = v7
        }
        v8 := GenS2{F0: 100}
        _ = fi5
        _ = v8
    }
    if (v1 != "hello") {
        v1 = ((v1 + ("foo bar" + (v1 + (v1 + (v1 + v1))))) + v1)
    }
    var v9 uint = 18446744073709551615
    if (-410 > 741) {
        var v10 uint16 = 1
        v1 = "a"
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{fi11, (^(fi11)), 9223372036854775807, 32767}
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := []int{127}
        v14 := GenS3{F0: -48, F1: 38, F2: 615}
        _ = v13
        _ = v14
    }
    if v3 {
        v15 := GenS4{F0: v3}
        _ = v15
    }
    if v3 {
        v3 = (109 < 159)
    } else {
        v16 := []int{-9223372036854775808}
        _ = v16
    }
    var v17 float32 = (((3.14159 * 3.14159) / 1.0) * float32(0.0))
    if (v17 >= -2.25) {
        var v18 uint = (v9 % (((v9 + 768)) | 1))
        v3 = (v9 == (v18 / ((4294967295) | 1)))
        _ = v18
    } else {
        var v19 int = 32768
        _ = v19
    }
    var v20 bool = (true || true)
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS2{F0: 95}
        v3 = true
        for fi23 := 0; fi23 < 8; fi23++ {
            fmt.Println(v9)
            if (242 != 52) {
                fmt.Println(v3, v20)
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v20 = (v9 >= 2)
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float64 = ((-2.25 - 0.0) - -2.25)
            v27 := GenS5{F0: 1, F1: 80, F2: 42}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    v9 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
