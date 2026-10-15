package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0() bool {
    var v6 float64 = -2.25
    var v7 bool = ("foo bar" >= "hello")
    switch ("" + (("" + "foo bar") + ("\n\t" + "\n\t"))) {
    case "foo bar":
        var v8 int32 = 1
        _ = v8
    case "a":
        v6 -= v6
    case "\n\t":
        v9 := GenS1{F0: 80, F1: 65535, F2: 128}
        _ = v9
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            v7 = v7
            _ = fi10
        }
    }
    _ = v6
    _ = v7
    return (18446744073709551615 <= 1)
}
func genFn1(p0 bool, p1 int64, p2 uint16) int32 {
    var v11 uint8 = 255
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch "a" {
            case "":
                v14 := []int{(-(fi12))}
                _ = v14
            }
            p1 -= (p1 | -27)
            if true {
                var v15 int8 = 100
                if p0 {
                    var v16 int = (fi13 ^ (fi12 ^ (fi12 << ((65535) & 63))))
                    v17 := []int{fi13, 32767, 517, fi12}
                    var v18 int8 = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v15
            } else {
                v19 := []int{(fi12 << ((4294967295) & 63)), (fi12 * fi13)}
                _ = v19
            }
            if p0 {
                var v20 int8 = 17
                var v21 int16 = 1
                _ = v20
                _ = v21
            } else {
                var v22 uint32 = uint32(v11)
                var v23 int = fi13
                var v24 int64 = p1
                _ = v22
                _ = v23
                _ = v24
            }
            v25 := []int{4294967295, fi13, 127, fi12}
            _ = fi13
            _ = v25
        }
        v26 := GenS1{F0: 69, F1: 2147483647, F2: v11}
        _ = fi12
        _ = v26
    }
    var v27 int64 = p1
    v11 -= (v11 ^ (v11 << ((614) & 63)))
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v27
    return -277
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.0 + (3.14159 + float64(1e10)))
    v3 := make(map[int]int)
    v4 := []uint32{2147483648, 65535, 518, 350}
    var v5 bool = false
    v28 := GenS1{F0: -1, F1: 1, F2: 128}
    for fi29 := 0; fi29 < 16; fi29++ {
        v30 := GenS1{F0: -128, F1: 124, F2: 0}
        v2 = (((v2 - float64(-2.25)) - 0.0) + v2)
        if ((("\n\t" + "") + ("hello" + "\n\t")) >= "\n\t") {
            for fi31 := 0; fi31 < 2; fi31++ {
                v5 = ((v5 || false) || (255 != 255))
                var v32 uint8 = 127
                v33 := GenS2{F0: (v5 && v5)}
                if v5 {
                    var v34 float32 = 1.5
                    _ = v34
                } else {
                    if (3.14159 > v2) {
                        if (!(((float32(v2) + (float32(1.5) + 3.14159)) - 1e10) >= 0.1)) {
                            v5 = (v2 > v2)
                        } else {
                            var v35 uint64 = 2
                            switch (fi29 * 1) {
                            case -911:
                                v5 = v5
                            }
                            v32 = v32
                            _ = v35
                        }
                        var v36 int16 = 559
                        _ = v36
                    }
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi29
        _ = v30
    }
    v2 -= (((v2 + -2.25) + v2) - v2)
    v5 = false
    v2 = (((v2 + ((1.5 + (v2 / 1.0)) + float64(float32(((((1e10 - float64((float32(v2) * 3.14159))) / float64((((float32((v2 * v2)) * 3.14159) + 0.0) / (((((((1e10 - -2.25) * 0.0) - 0.1) - 1e10) - (((1e10 / (float32(v2) / float32(1.5))) - (-2.25 + (0.0 - 0.1))) * (3.14159 * -2.25))) * 1.5) - 0.0)))) - v2) + v2))))) / v2) / v2)
    switch 42 {
    case 32767:
        for fi37 := 0; fi37 < 11; fi37++ {
            if v5 {
                switch (("foo bar" + "") + ("\n\t" + ("" + "a"))) {
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    var v38 uint8 = 37
                    _ = v38
                case "hello":
                    v5 = (true || (!(((v2 + v2) == v2) || v5)))
                }
                v39 := GenS2{F0: false}
                var v40 int32 = 0
                _ = v39
                _ = v40
            }
            var v41 uint16 = 32767
            v41 -= (v41 * 2)
            _ = fi37
            _ = v41
        }
    }
    if (18446744073709551615 != 4294967295) {
        v2 += ((v2 / (v2 * (1.5 * (((v2 * 0.0) / (float64(-2.25) - v2)) * 1e10)))) + ((3.14159 / float64((-2.25 * (float32(3.14159) - (1e10 + 1e10))))) / float64(-2.25)))
        var v42 bool = (1 == -128)
        _ = v42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        if (42 >= 128) {
            v5 = v5
            var v44 float64 = (float64(((1e10 + 1e10) - (((0.1 + 1e10) - float32((3.14159 + v2))) * float32(v2)))) + (0.1 * v2))
            _ = v44
        }
        _ = fi43
    }
    if v5 {
        fmt.Println(v2)
        fmt.Println(v5, v5)
    } else {
        if (33 >= 255) {
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 uint = 622
                var v47 bool = v5
                _ = fi45
                _ = v46
                _ = v47
            }
            var v48 int32 = -2147483648
            if v5 {
                for fi49 := 0; fi49 < 11; fi49++ {
                    var v50 string = "hello"
                    v5 = (true && v5)
                    fmt.Println(v5)
                    _ = fi49
                    _ = v50
                }
            } else {
                v51 := GenS3{F0: 745}
                fmt.Println(v2)
                _ = len(v4)
                _ = v51
            }
            _ = v48
        }
    }
    switch "\n\t" {
    case "foo bar":
        switch "" {
        case "a":
            fmt.Println(v2, v5)
        default:
            if (!v5) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi52 := 0; fi52 < 2; fi52++ {
                    v2 = (0.0 + float64(((3.14159 * ((3.14159 / 1e10) + 0.1)) / 1.5)))
                    var v53 string = ("\n\t" + ("foo bar" + "hello"))
                    _ = fi52
                    _ = v53
                }
                v54 := GenS4{F0: 1, F1: false}
                _ = v54
            } else {
                v55 := GenS3{F0: 65535}
                v2 -= ((3.14159 * v2) - v2)
                switch -803 {
                case 557:
                    var v56 int16 = -1
                    _ = v56
                case 9223372036854775807:
                    var v57 uint16 = 0
                    _ = v57
                case 1:
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                default:
                    var v58 uint8 = 176
                    _ = v58
                }
                _ = v55
            }
        }
    case "a":
        _ = len(v3)
    case "\n\t":
        fmt.Println(v2)
    }
    fmt.Println(v2, v5)
    v2 = float64(0.1)
    if (v5 && v5) {
        fmt.Println(v5)
        for fi59 := 0; fi59 < 10; fi59++ {
            var v60 int8 = 42
            _ = fi59
            _ = v60
        }
    }
    if (((1e10 / 1.0) != (3.14159 - 0.0)) || (((!(1000 != -863)) || v5) && true)) {
        var v61 float32 = -2.25
        switch -1 {
        case -685:
            v61 -= 0.0
        default:
            if v5 {
                var v62 uint32 = 2
                _ = v62
            }
        }
        if v5 {
            if true {
                v2 = v2
            }
            var v63 uint64 = 42
            _ = v63
        } else {
            for fi64 := 0; fi64 < 12; fi64++ {
                var v65 int16 = -328
                _ = fi64
                _ = v65
            }
            v5 = v5
            delete(v3, 59)
        }
        _ = v61
    } else {
        fmt.Println(v5, v5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if v5 {
        for fi66 := 0; fi66 < 20; fi66++ {
            switch ((("\n\t" + "\n\t") + "foo bar") + "a") {
            case "hello":
                v2 += v2
            default:
                if (-1 == 1) {
                    _ = len(v4)
                    var v67 string = "hello"
                    _ = v67
                }
            }
            v68 := GenS5{F0: 1}
            v69 := GenS1{F0: 99, F1: 2147483647, F2: 130}
            _ = fi66
            _ = v68
            _ = v69
        }
        v70 := GenS2{F0: false}
        _ = v70
    } else {
        v2 += float64(1e10)
    }
    var v71 int32 = 661
    v5 = ((v71 << ((741) & 63)) > (^(v71)))
    switch ("\n\t" + "") {
    case "foo bar":
        for fi72 := 0; fi72 < 16; fi72++ {
            var v73 string = "foo bar"
            _ = fi72
            _ = v73
        }
    case "":
        fmt.Println(v71, v71)
    default:
        v71 -= v71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
