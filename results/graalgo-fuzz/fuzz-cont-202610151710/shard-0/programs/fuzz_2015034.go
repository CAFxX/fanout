package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = -1
    var v3 string = v1
    var v4 uint8 = 255
    var v5 float64 = float64(1.5)
    if true {
        switch 32768 {
        case 127:
            switch 42 {
            case 2147483647:
                switch 9223372036854775807 {
                case -1:
                    v6 := GenS1{F0: true, F1: 42}
                    _ = v6
                case 0:
                    var v7 int32 = -2147483648
                    _ = v7
                default:
                    var v8 uint = 2
                    _ = v8
                }
            case 42:
                v4 -= 128
            default:
                v2 = int64(v4)
            }
        case 42:
            if (-2.25 < 1.5) {
                if false {
                    v9 := []int{}
                    v10 := GenS1{F0: false, F1: 165}
                    if (v4 >= 2) {
                        fmt.Println(v1)
                    }
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{}
                    if len(v11) > 0 {
                        v11[0] = 655
                    }
                    v12 := GenS1{F0: true, F1: 42}
                    _ = v11
                    _ = v12
                }
                fmt.Println(v5, v2)
                var v13 uint32 = 0
                _ = v13
            } else {
                if false {
                    var v14 uint = 65535
                    v15 := GenS1{F0: (-2.25 <= (1.5 + (0.1 / (0.0 / (3.14159 + ((3.14159 - (float32((v5 / v5)) * 3.14159)) / 1e10)))))), F1: 0}
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: (1000 == -656), F1: 922}
                    _ = v16
                }
                var v17 bool = (0.0 < 3.14159)
                _ = v17
            }
        case 32768:
            var v18 int64 = v2
            _ = v18
        default:
            var v19 int = 0
            _ = v19
        }
        v3 = (((("hello" + ((("\n\t" + (v3 + ((v3 + "\n\t") + "hello"))) + v3) + v3)) + v3) + "\n\t") + ("\n\t" + (v1 + (v3 + v3))))
        v20 := []int{77, 2147483647}
        _ = v20
    }
    var v21 string = ("foo bar" + (("hello" + "foo bar") + v1))
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            switch v1 {
            case "hello":
                if false {
                    var v24 uint64 = 18446744073709551615
                    v25 := []int{fi22}
                    if (((-42 != 1) || (1 <= 42)) && false) {
                        v25 = append(v25, (-(fi22)))
                        v21 = ((v21 + v3) + v3)
                        switch v1 {
                        case "a":
                            var v26 uint = 1
                            _ = v26
                        default:
                            if false {
                                if false {
                                    v4 -= (v4 & (v4 | (v4 + v4)))
                                }
                            } else {
                                if len(v25) > 0 {
                                    v25[0] = fi23
                                }
                                v4 -= v4
                            }
                        }
                    } else {
                        fmt.Println(v3, fi22)
                    }
                    v3 = v3
                    if true {
                        if ((3.14159 + 1.5) != ((float32((((v5 / (((v5 - v5) + (v5 * 3.14159)) * v5)) + float64(1e10)) - (3.14159 * 3.14159))) + 0.1) + 1e10)) {
                            v27 := GenS2{F0: -36, F1: -509}
                            fmt.Println(fi23)
                            switch v21 {
                            case "\n\t":
                                var v28 int16 = 32767
                                _ = v28
                            default:
                                fmt.Println(fi22, fi22)
                            }
                            _ = v27
                        }
                        _ = len(v25)
                        var v29 uint16 = uint16(v4)
                        _ = v29
                    }
                    _ = v24
                    _ = v25
                } else {
                    v5 += 0.1
                    switch fi23 {
                    case 32768:
                        if true {
                            var v30 int8 = -1
                            v31 := []int{9223372036854775807}
                            v2 -= (v2 >> ((255) & 63))
                            v1 = v3
                            v2 = -1
                            _ = v30
                            _ = v31
                        }
                    case 32767:
                        v21 = (v21 + "\n\t")
                    default:
                        switch (fi23 % (((fi22 ^ (fi23 & -1))) | 1)) {
                        case 9223372036854775807:
                            v2 = v2
                            v32 := []int{(-(fi22)), (fi23 << ((65535) & 63)), (fi22 - (fi22 + (fi22 << ((125) & 63)))), -330}
                            _ = v32
                        default:
                            fmt.Println(v5)
                        }
                    }
                }
                var v33 uint = 2
                _ = v33
            case "a":
                fmt.Println(v1)
            }
            if false {
                v4 -= 127
                v34 := []int{(fi23 | (fi23 / (((fi23 << ((2) & 63))) | 1)))}
                var v35 uint16 = 32767
                _ = v34
                _ = v35
            }
            v36 := []int{fi23, fi23}
            var v37 string = ((((((v1 + (v3 + ((v1 + (v21 + "\n\t")) + (v21 + v1)))) + "foo bar") + ((("a" + (v3 + v3)) + ((((v21 + v3) + "\n\t") + "hello") + ("hello" + "a"))) + v21)) + "foo bar") + ("a" + v3)) + v21)
            _ = fi23
            _ = v36
            _ = v37
        }
        v3 = v21
        _ = fi22
    }
    switch -314 {
    case 2147483647:
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 uint32 = 42
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        v4 = v4
        var v41 int64 = (v2 - v2)
        v41 = v41
        _ = fi40
        _ = v41
    }
    var v42 uint16 = 1
    var v43 float64 = (v5 - v5)
    fmt.Println(v21, v5, v42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
