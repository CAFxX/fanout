package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0() int8 {
    switch (("a" + ("a" + "a")) + "hello") {
    case "a":
        v7 := []int{159, 127, 32767}
        _ = v7
    case "hello":
        var v8 bool = ((209 == 126) || (2147483648 < 1))
        _ = v8
    case "":
        v9 := []int{}
        _ = v9
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 bool = (255 != 42)
            switch fi10 {
            case -2147483648:
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 float32 = float32(0.0)
                    _ = fi12
                    _ = v13
                }
            case 128:
                fmt.Println(fi10)
            default:
                switch fi10 {
                case 820:
                    var v14 string = "a"
                    _ = v14
                case 4294967295:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        fmt.Println(v11)
                        _ = fi15
                    }
                case -9223372036854775808:
                    if (false && ((!(42 > 960)) && v11)) {
                        fmt.Println(v11)
                        v16 := GenS1{F0: -9223372036854775808, F1: 687, F2: v11}
                        _ = v16
                    } else {
                        v11 = ((v11 && (2 > 255)) && (1e10 > float32(0.0)))
                        if true {
                            if ((!(!v11)) || v11) {
                                if v11 {
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        var v18 int64 = -662
                                        var v19 uint = 1
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                    }
                                    var v20 uint16 = 253
                                    v21 := []int{(fi10 / ((fi10) | 1)), (-(fi10)), 128}
                                    _ = v20
                                    _ = v21
                                }
                                v22 := GenS2{F0: 185, F1: 1, F2: -2147483648}
                                var v23 int32 = -2147483648
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint8 = 128
                                v24 += 255
                                _ = v24
                            }
                        }
                    }
                }
            }
            var v25 uint = 42
            _ = fi10
            _ = v11
            _ = v25
        }
    }
    switch -654 {
    case 9223372036854775807:
        var v26 uint8 = 1
        _ = v26
    case 0:
        for fi27 := 0; fi27 < 19; fi27++ {
            switch ("hello" + "foo bar") {
            case "a":
                v28 := GenS1{F0: -13, F1: 530, F2: (true && (true || true))}
                _ = v28
            case "hello":
                if true {
                    v29 := []int{0, fi27, (fi27 % (((fi27 - (fi27 >> ((0) & 63)))) | 1)), fi27}
                    switch fi27 {
                    case 79:
                        var v30 string = ("" + ((("" + (("foo bar" + "hello") + "a")) + ((("\n\t" + "a") + "foo bar") + "a")) + "hello"))
                        _ = v30
                    }
                    if true {
                        var v31 int = fi27
                        _ = v31
                    }
                    _ = v29
                } else {
                    var v32 float32 = (-2.25 * 1e10)
                    v33 := GenS3{F0: fi27, F1: 2, F2: 512}
                    _ = v32
                    _ = v33
                }
            case "\n\t":
                for fi34 := 0; fi34 < 18; fi34++ {
                    if true {
                        var v35 uint32 = 0
                        v36 := []int{(fi27 - (fi27 % ((fi34) | 1)))}
                        fmt.Println(fi34, v35, fi34)
                        v35 += (v35 % (((v35 + (v35 & (v35 & v35)))) | 1))
                        v37 := GenS3{F0: fi27, F1: 670, F2: 0}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    _ = fi34
                }
            default:
                var v38 int32 = -2147483648
                _ = v38
            }
            switch 2147483647 {
            case -2147483648:
                var v39 string = "a"
                _ = v39
            }
            v40 := []int{fi27, fi27, (fi27 ^ 9223372036854775807), (fi27 * fi27)}
            _ = fi27
            _ = v40
        }
    case -514:
        var v41 int32 = 0
        _ = v41
    default:
        v42 := []int{999, 128, 4294967295}
        _ = v42
    }
    var v43 string = ""
    _ = v43
    return 121
}

func main() {
    var v1 int32 = 1
    v2 := []int{-9223372036854775808, 32768}
    var v3 float32 = 1e10
    var v4 uint32 = 0
    var v5 int8 = 127
    var v6 int32 = (v1 ^ 0)
    _ = len(v2)
    var v44 int16 = -1
    var v45 uint64 = 1
    v46 := GenS4{F0: 2147483648, F1: false, F2: 586}
    var v47 uint8 = 1
    var v48 int32 = v1
    var v49 uint8 = (v47 << ((18446744073709551615) & 63))
    var v50 string = ""
    if len(v2) > 0 {
        v2[0] = 0
    }
    v47 -= (v49 - 2)
    switch 0 {
    case 32768:
        fmt.Println(v50, v45, v50)
    default:
        switch 127 {
        case 127:
            v45 -= (v45 & v45)
        case -81:
            fmt.Println(v3, v44, v44)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
