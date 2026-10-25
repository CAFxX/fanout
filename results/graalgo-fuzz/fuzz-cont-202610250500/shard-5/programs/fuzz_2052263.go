package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 uint64) string {
    p0 = (p0 - p0)
    var v5 uint32 = 2147483648
    v5 += (v5 / (((v5 << ((72) & 63))) | 1))
    _ = p0
    _ = v5
    return (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))
}
func genFn1(p0 uint8, p1 uint) int {
    v6 := []int{128, 278, 0, -9223372036854775808}
    if false {
        fmt.Println(p1, p1)
        for fi7 := 0; fi7 < 6; fi7++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p0 = (p0 * (p0 + 1))
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 303
}

func main() {
    var v1 uint = 811
    var v2 uint = (v1 - 255)
    var v3 int16 = 0
    v4 := make(map[int]string)
    v3 += (v3 >> ((v1) & 63))
    for fi8 := 0; fi8 < 10; fi8++ {
        v1 += (v1 / (((v1 - (v2 << ((326) & 63)))) | 1))
        _ = fi8
    }
    fmt.Println(v1, v3, v3)
    var v9 uint = (v2 ^ v2)
    v9 = v9
    switch "\n\t" {
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            if true {
                var v11 bool = false
                switch "a" {
                case "":
                    if v11 {
                        v12 := GenS1{F0: -1, F1: uint8(v2), F2: (-(v3))}
                        _ = v12
                    }
                case "a":
                    var v13 int8 = 100
                    _ = v13
                case "hello":
                    for fi14 := 0; fi14 < 7; fi14++ {
                        fmt.Println(v1)
                        if false {
                            v11 = true
                            var v15 string = ("foo bar" + "foo bar")
                            v16 := []int{(fi10 / (((fi14 ^ fi10)) | 1))}
                            _ = v15
                            _ = v16
                        }
                        _ = fi14
                    }
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 float32 = -2.25
                    _ = fi17
                    _ = v18
                }
                _ = v11
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int8 = 42
                    _ = fi19
                    _ = v20
                }
                v2 -= (v9 + 42)
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                delete(v4, genFn1(1, 18446744073709551615))
                v22 := GenS1{F0: 0, F1: 34, F2: 1}
                v23 := GenS2{F0: v2, F1: 1, F2: 717}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi10
        }
    case "":
        fmt.Println(v2)
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int64 = 821
        if (false || ((!false) && (1.5 == 3.14159))) {
            var v26 int = (fi24 - fi24)
            _ = v26
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint = 2
                _ = fi27
                _ = v28
            }
        }
        _ = fi24
        _ = v25
    }
    v3 = -863
    v29 := []int{1}
    v9 = (v1 + (v1 >> ((42) & 63)))
    v30 := GenS3{F0: 255}
    _ = len(v29)
    v9 = v2
    if ((("\n\t" + "a") + "hello") <= "a") {
        v31 := GenS4{F0: false, F1: (!(100 < 4294967295)), F2: "hello"}
        switch ("" + "\n\t") {
        case "\n\t":
            var v32 int64 = 0
            _ = v32
        case "foo bar":
            v29 = append(v29, 4294967295)
        }
        _ = v31
    }
    v3 += v3
    v1 = (v9 ^ (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
