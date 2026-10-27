package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int8 = 1
    var v2 uint32 = 2
    var v3 float32 = float32(0.1)
    var v4 float64 = (((1.5 + float64(float32(1e10))) * (0.0 / 1e10)) * float64(-2.25))
    var v5 uint16 = 32767
    var v6 int = 1
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v8 := GenS1{F0: 128, F1: fi7}
            v2 -= v2
            switch (fi7 + (-(fi7))) {
            case 1:
                if true {
                    v9 := []int{v6}
                    _ = v9
                } else {
                    v10 := GenS1{F0: 0, F1: 42}
                    _ = v10
                }
            }
            _ = v8
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS2{F0: "hello", F1: 42}
            if true {
                v1 = (v1 >> ((18446744073709551615) & 63))
                v13 := []int{fi7, v6}
                var v14 float32 = v3
                _ = v13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    var v15 uint8 = 128
    v5 = v5
    v16 := []int{v6, (v6 * v6), -8, (v6 / ((v6) | 1))}
    v16 = append(v16, -649)
    v17 := GenS3{F0: 1, F1: (^(v15))}
    _ = len(v16)
    v4 = v4
    v6 -= (v6 ^ (v6 | v6))
    switch (v6 + -756) {
    case 32767:
        if ((("a" + "") + ("a" + ("foo bar" + ("foo bar" + ("a" + "a"))))) == "a") {
            if len(v16) > 0 {
                v16[0] = (v6 * 305)
            }
            if true {
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 uint32 = (v2 >> ((255) & 63))
                    _ = fi18
                    _ = v19
                }
            }
            v6 += (v6 - 0)
        }
    case 4294967295:
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                v2 = (v2 - v2)
                _ = fi21
            }
            _ = fi20
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v2 = (v2 - 1)
        _ = fi22
    }
    var v23 float32 = v3
    switch "hello" {
    case "a":
        _ = len(v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
