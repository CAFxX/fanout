package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 int32) int8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 += p0
        switch ("a" + "") {
        case "foo bar":
            v6 := []int{1, -679}
            _ = v6
        case "":
            switch int(p0) {
            case 116:
                p0 = p0
            case 593:
                v7 := GenS1{F0: p0, F1: 0}
                _ = v7
            case 1:
                p0 = 1
            }
        }
        _ = fi5
    }
    v8 := GenS2{F0: "hello"}
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 float32 = (-2.25 + (1.5 * (float32(float64((float32(1.5) / float32((-2.25 + (3.14159 + (1e10 / (1e10 + (((float64(3.14159) - (3.14159 + (0.0 / 3.14159))) - 0.0) - 3.14159))))))))) / 1.0)))
    var v2 bool = true
    var v3 bool = (555 >= 0)
    var v4 int = 42
    var v9 int32 = -630
    switch ("hello" + "hello") {
    case "hello":
        v10 := GenS1{F0: (-(v9)), F1: 18446744073709551615}
        _ = v10
    }
    var v11 uint32 = 120
    v12 := GenS3{F0: 2, F1: 9223372036854775807}
    var v13 int8 = -119
    var v14 uint64 = 225
    var v15 uint32 = v11
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            if ((v15 / (((v15 / ((v11) | 1))) | 1)) <= 1) {
                v18 := []int{fi16, v4, (-(fi17))}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                v9 -= v9
                v19 := GenS4{F0: v15, F1: 0, F2: 42}
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        v20 := []int{(v4 << ((42) & 63)), v4, 42, fi16}
        _ = fi16
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        if v3 {
            for fi22 := 0; fi22 < 12; fi22++ {
                v4 -= (v4 / ((fi21) | 1))
                _ = fi22
            }
        }
        var v23 uint16 = 65535
        switch (fi21 / ((fi21) | 1)) {
        case 109:
            var v24 int32 = (^(v9))
            _ = v24
        case -446:
            v25 := []int{fi21, fi21, (v4 | int(v14)), (fi21 >> ((1) & 63))}
            _ = v25
        case 32767:
            switch 32768 {
            case 42:
                var v26 uint32 = v11
                _ = v26
            }
        default:
            switch (("\n\t" + (("foo bar" + "foo bar") + "foo bar")) + ("" + "a")) {
            case "hello":
                var v27 string = "hello"
                _ = v27
            case "":
                v15 = (v11 >> ((uint(v23)) & 63))
            }
        }
        _ = fi21
        _ = v23
    }
    var v28 uint8 = 2
    v9 += v9
    var v29 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
