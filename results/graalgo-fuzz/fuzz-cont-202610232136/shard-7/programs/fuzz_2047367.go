package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint64 }
func genFn0() uint8 {
    var v7 int = 32767
    v8 := []int{v7, (-(v7)), 127}
    v7 += 128
    for fi9 := 0; fi9 < 11; fi9++ {
        v7 = v7
        _ = fi9
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    var v6 bool = (((-1 == 0) || ((false || (51 == -125)) && (1 <= 168))) && (float32((1.5 - (-2.25 - 0.1))) == (1.5 / (0.0 - float32(1.5)))))
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(fi11, fi11, fi10)
            delete(v2, (fi10 ^ fi11))
            _ = fi11
        }
        if v6 {
            var v12 float32 = 3.14159
            var v13 float32 = v12
            var v14 uint = 4294967295
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := GenS1{F0: (fi10 - 0)}
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint32 = 42
                v6 = false
                switch "hello" {
                case "\n\t":
                    var v18 uint64 = 336
                    if ((!((v6 && (v18 <= 42)) || true)) || ((!v6) && true)) {
                        fmt.Println(v17, fi16)
                        if v6 {
                            v6 = (-119 >= 127)
                            var v19 string = (("" + ("foo bar" + "a")) + ("a" + "hello"))
                            v18 = (v18 + (v18 * (^(v18))))
                            _ = v19
                        } else {
                            fmt.Println(fi10, v17)
                            var v20 uint = uint(v17)
                            _ = v20
                        }
                        var v21 bool = ((32767 <= 2) && true)
                        _ = v21
                    } else {
                        v22 := GenS1{F0: fi10}
                        switch (("\n\t" + ("foo bar" + (((("" + "hello") + ("" + (("" + "\n\t") + ("a" + "\n\t")))) + "") + "a"))) + "a") {
                        case "a":
                            v23 := GenS1{F0: (fi16 - (fi10 >> ((35) & 63)))}
                            _ = v23
                        case "":
                            v24 := GenS1{F0: fi10}
                            _ = v24
                        case "hello":
                            v25 := []int{fi16, (^(fi10)), (fi16 >> ((255) & 63)), fi16}
                            switch 0 {
                            case -1:
                                var v26 uint8 = 0
                                _ = v26
                            case 32768:
                                v27 := GenS2{F0: fi10, F1: -955, F2: 4294967295}
                                _ = v27
                            case -2147483648:
                                v28 := GenS2{F0: fi16, F1: -599, F2: 231}
                                _ = v28
                            default:
                                var v29 int32 = 37
                                _ = v29
                            }
                            _ = v25
                        default:
                            v18 = v18
                        }
                        fmt.Println(fi16, v17, fi10)
                        fmt.Println(v18, fi16, v6)
                        if (v6 && (0.1 == 3.14159)) {
                            var v30 float32 = float32(float64((-2.25 + 1.5)))
                            _ = v30
                        } else {
                            v31 := GenS1{F0: (fi16 * -958)}
                            switch ("hello" + ((("hello" + ("\n\t" + "")) + "\n\t") + "a")) {
                            case "\n\t":
                                var v32 int32 = -2147483648
                                _ = v32
                            case "hello":
                                var v33 int64 = -730
                                _ = v33
                            case "":
                                v34 := GenS3{F0: "hello"}
                                _ = v34
                            }
                            fmt.Println(v6)
                            _ = v31
                        }
                        _ = v22
                    }
                    _ = v18
                case "hello":
                    if false {
                        v35 := []int{fi16, 127}
                        delete(v2, (fi16 >> ((1) & 63)))
                        v17 -= v17
                        _ = v35
                    }
                default:
                    var v36 int64 = -573
                    _ = v36
                }
                _ = fi16
                _ = v17
            }
            v6 = (false || false)
            _ = v15
        }
        v6 = v6
        _ = fi10
    }
    v6 = (!((!v6) || v6))
    v6 = true
    v37 := GenS4{F0: -805, F1: -87, F2: 127}
    fmt.Println(v6)
    v3[921] = (42 > 100)
    v38 := GenS5{F0: 514, F1: 4294967295}
    v6 = false
    var v39 int8 = 42
    v6 = v6
    v40 := GenS6{F0: 0, F1: 2, F2: 956}
    delete(v4, 0)
    var v41 int8 = -128
    v41 = 0
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
