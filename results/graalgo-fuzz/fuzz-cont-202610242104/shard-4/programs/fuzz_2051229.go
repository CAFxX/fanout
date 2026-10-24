package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0(p0 uint32) int16 {
    v7 := []int{-428, 32768, 2147483647}
    v8 := GenS1{F0: 199, F1: 2147483647, F2: -338}
    var v9 int64 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -32768
}
func genFn1(p0 uint16, p1 uint8) int16 {
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float64 = 1e10
        v12 := []int{fi10}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int32 = -1
    fmt.Println(p1, v13, p0)
    fmt.Println(p0, p0, p0)
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 4
    var v3 uint32 = 765
    var v4 bool = true
    var v5 string = ("\n\t" + "\n\t")
    var v6 float32 = 0.1
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 bool = true
        _ = fi15
        _ = v16
    }
    v5 = (("hello" + "a") + ("foo bar" + v5))
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 string = (v5 + ((v5 + ((v5 + "hello") + v5)) + v5))
        v18 = v18
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v3, v5)
    switch ("\n\t" + "") {
    case "":
        if v4 {
            v4 = (int(v3) != 9223372036854775807)
        }
    case "foo bar":
        var v19 bool = v4
        _ = v19
    case "\n\t":
        fmt.Println(v2, v2, v4)
    }
    var v20 bool = (42 <= 255)
    if v20 {
        if (v4 && v4) {
            if v4 {
                v3 = 1
                v2 -= v2
                if true {
                    if true {
                        var v21 bool = (!true)
                        var v22 float64 = (1.5 / (1.5 * (0.1 - ((1.5 / float64(v6)) + 3.14159))))
                        _ = len(v1)
                        _ = v21
                        _ = v22
                    }
                    v23 := v1[42]
                    _ = v23
                } else {
                    delete(v1, 127)
                    v5 = v5
                }
            }
        } else {
            v4 = (v20 && v20)
            v24 := GenS2{F0: 808, F1: 32767, F2: v5}
            v2 = (v2 * int8(v3))
            _ = v24
        }
        v25 := []int{0}
        v26 := v1[-343]
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            v20 = true
            _ = fi27
        }
        v28 := GenS1{F0: 128, F1: -748, F2: 42}
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        fmt.Println(v4)
        var v30 string = ("" + v5)
        var v31 uint32 = (v3 + v3)
        _ = fi29
        _ = v30
        _ = v31
    }
    if (!(true || v4)) {
        if (false || true) {
            var v32 float32 = v6
            v6 -= v32
            _ = v32
        }
        fmt.Println(v4, v20, v6)
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 int16 = -1
            v35 := GenS1{F0: 1, F1: -1, F2: 42}
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 5; fi36++ {
            v6 = (1.5 / 1e10)
            v3 = 2147483647
            _ = fi36
        }
    }
    if (!v20) {
        for fi37 := 0; fi37 < 3; fi37++ {
            for fi38 := 0; fi38 < 7; fi38++ {
                v39 := GenS2{F0: 0, F1: 2147483647, F2: "a"}
                v40 := GenS3{F0: 0, F1: 9223372036854775807, F2: 0}
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi37
        }
    }
    if ((v3 ^ v3) >= (v3 & (v3 - v3))) {
        for fi41 := 0; fi41 < 13; fi41++ {
            v42 := GenS4{F0: (833 <= 1), F1: 775}
            _ = fi41
            _ = v42
        }
        if (605 != 810) {
            if v20 {
                if false {
                    _ = len(v1)
                }
            } else {
                for fi43 := 0; fi43 < 8; fi43++ {
                    var v44 uint64 = 18446744073709551615
                    switch (fi43 / (((fi43 >> ((4294967295) & 63))) | 1)) {
                    case 1:
                        v2 = v2
                    }
                    _ = fi43
                    _ = v44
                }
                var v45 int8 = v2
                var v46 uint32 = (^(v3))
                _ = v45
                _ = v46
            }
            for fi47 := 0; fi47 < 16; fi47++ {
                var v48 int8 = 89
                _ = fi47
                _ = v48
            }
        } else {
            fmt.Println(v3, v2)
        }
    }
    var v49 int32 = -757
    if v4 {
        if ((v3 + 65535) > (v3 % ((v3) | 1))) {
            if (42 >= 2) {
                v50 := []int{}
                fmt.Println(v49, v20)
                _ = len(v50)
                _ = v50
            } else {
                if v20 {
                    var v51 uint64 = 18446744073709551615
                    var v52 uint = 437
                    _ = v51
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 13; fi53++ {
                        v20 = (v5 != "a")
                        for fi54 := 0; fi54 < 11; fi54++ {
                            var v55 int32 = (v49 << ((4294967295) & 63))
                            var v56 uint16 = 119
                            v57 := GenS2{F0: 18446744073709551615, F1: (fi53 >> ((321) & 63)), F2: v5}
                            if ((65535 <= (v56 << ((1) & 63))) && false) {
                                v58 := GenS3{F0: (^(v56)), F1: 9223372036854775807, F2: uint8(v56)}
                                _ = v58
                            } else {
                                fmt.Println(v5, v49)
                                v1[fi54] = fi54
                            }
                            _ = fi54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        if true {
                            for fi59 := 0; fi59 < 7; fi59++ {
                                fmt.Println(fi53, v20, v5)
                                v49 = 42
                                delete(v1, (^(fi53)))
                                _ = fi59
                            }
                        } else {
                            v60 := GenS1{F0: 128, F1: 896, F2: v49}
                            _ = v60
                        }
                        _ = fi53
                    }
                    switch -2147483648 {
                    case 0:
                        fmt.Println(v5)
                    case 4294967295:
                        v5 = v5
                    default:
                        fmt.Println(v20, v4, v6)
                    }
                }
                if true {
                    v61 := GenS4{F0: false, F1: 255}
                    fmt.Println(v5)
                    _ = v61
                } else {
                    v4 = false
                }
                switch 9223372036854775807 {
                case 32768:
                    _ = len(v1)
                case -248:
                    delete(v1, 1)
                default:
                    var v62 int = 32768
                    _ = v62
                }
            }
        } else {
            for fi63 := 0; fi63 < 3; fi63++ {
                if (!v4) {
                    v3 -= (v3 - v3)
                }
                _ = fi63
            }
            var v64 uint8 = 243
            v65 := []int{}
            _ = v64
            _ = v65
        }
        v66 := GenS5{F0: 0, F1: v3, F2: (0 >= 643)}
        v67 := []int{int(v3), 2147483647}
        _ = v66
        _ = v67
    }
    var v68 float32 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v68
}
