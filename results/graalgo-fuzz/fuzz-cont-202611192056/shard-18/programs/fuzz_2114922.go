package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 int64 = 0
    var v4 int64 = v3
    var v5 float64 = 1.5
    switch -9223372036854775808 {
    case -1:
        if ((!("hello" != "")) || (!false)) {
            v5 -= v5
            if (42 >= v1) {
                var v6 float32 = 1e10
                _ = v6
            }
        } else {
            switch ("foo bar" + "a") {
            case "":
                v2[(^(v1))] = "a"
            }
            for fi7 := 0; fi7 < 6; fi7++ {
                v5 += float64(0.0)
                switch (-(fi7)) {
                case 32768:
                    v8 := GenS1{F0: 18446744073709551615, F1: 65535, F2: v1}
                    _ = v8
                case -708:
                    var v9 string = (("" + "hello") + "")
                    _ = v9
                case -1:
                    if (-672 == 1) {
                        for fi10 := 0; fi10 < 2; fi10++ {
                            if ("a" > "") {
                                var v11 int32 = 276
                                var v12 string = ""
                                if true {
                                    v3 -= v4
                                }
                                _ = v11
                                _ = v12
                            } else {
                                v13 := []int{42}
                                var v14 uint16 = 988
                                v5 = 3.14159
                                var v15 int8 = 19
                                v3 = -9223372036854775808
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            _ = fi10
                        }
                        fmt.Println(v3, fi7, v1)
                    }
                default:
                    fmt.Println(v4, fi7)
                }
                _ = fi7
            }
        }
    case 2147483647:
        v16 := GenS1{F0: 475, F1: 32767, F2: 239}
        _ = v16
    case 32768:
        if false {
            v1 += v1
        }
    }
    v4 = v3
    fmt.Println(v5)
    fmt.Println(v5, v3)
    if ((v4 ^ v4) > -160) {
        v5 -= float64(1e10)
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(v1, v5)
                v4 = 1
                fmt.Println(fi18, fi18)
                v3 += (v4 | v3)
                var v19 uint32 = v1
                _ = fi18
                _ = v19
            }
            if true {
                v2[v1] = "foo bar"
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS2{F0: v1}
                _ = fi20
                _ = v21
            }
            _ = fi17
        }
        v22 := []int{-546, -9223372036854775808, 815}
        _ = v22
    }
    fmt.Println(v1, v3)
    v3 -= (v4 | v4)
    var v23 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
