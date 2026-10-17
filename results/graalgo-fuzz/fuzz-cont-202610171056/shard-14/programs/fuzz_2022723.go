package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 bool) int8 {
    p0 = (1 == 794)
    p0 = p0
    _ = p0
    return 1
}
func genFn1() string {
    if true {
        var v7 uint16 = 850
        if false {
            var v8 uint = 4294967295
            v7 = (v7 & (v7 | 742))
            _ = v8
        } else {
            v9 := []int{-618, 1, 591, -235}
            for fi10 := 0; fi10 < 9; fi10++ {
                v7 = v7
                if true {
                    v11 := GenS1{F0: v7, F1: "hello", F2: 1}
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := GenS1{F0: v7, F1: (("\n\t" + "\n\t") + "foo bar"), F2: 2}
                        if (!(1e10 > 3.14159)) {
                            _ = len(v9)
                            if len(v9) > 0 {
                                v9[0] = -736
                            }
                            if (122 != 140) {
                                switch -2147483648 {
                                case 32768:
                                    var v14 uint64 = 796
                                    _ = v14
                                case 464:
                                    v15 := GenS2{F0: v7, F1: -2147483648}
                                    _ = v15
                                default:
                                    fmt.Println(fi12)
                                }
                                var v16 int16 = 0
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = v16
                            }
                        } else {
                            fmt.Println(fi10, fi12, v7)
                            v17 := GenS1{F0: (v7 ^ 32767), F1: "foo bar", F2: 1}
                            var v18 float32 = 1.5
                            _ = v17
                            _ = v18
                        }
                        _ = fi12
                        _ = v13
                    }
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch fi10 {
                        case 32768:
                            v9 = append(v9, (fi19 >> ((267) & 63)))
                        case 32767:
                            v7 = (v7 << ((65535) & 63))
                        }
                        _ = fi19
                    }
                    _ = v11
                }
                var v20 string = "\n\t"
                _ = fi10
                _ = v20
            }
            if false {
                v7 -= 112
            } else {
                if true {
                    v7 = v7
                } else {
                    v21 := GenS3{F0: 426}
                    for fi22 := 0; fi22 < 12; fi22++ {
                        _ = len(v9)
                        v7 -= (v7 / (((v7 >> ((4294967295) & 63))) | 1))
                        _ = fi22
                    }
                    v7 -= (v7 | v7)
                    _ = v21
                }
                v7 += v7
                v7 += v7
            }
            _ = v9
        }
        _ = v7
    }
    var v23 int16 = 1000
    v23 = v23
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(fi24)
        v25 := []int{-804, 4294967295}
        if (genFn0((978 >= 1)) <= int8(fi24)) {
            if (3.14159 > float32(3.14159)) {
                if (127 > 42) {
                    if (64 != 255) {
                        v23 = v23
                    }
                } else {
                    v26 := GenS4{F0: int16(fi24), F1: 0, F2: "foo bar"}
                    var v27 int = fi24
                    _ = v26
                    _ = v27
                }
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                var v28 uint8 = 255
                _ = v28
            } else {
                if (!true) {
                    v25 = append(v25, 0)
                    v29 := GenS4{F0: v23, F1: 128, F2: "a"}
                    _ = v29
                } else {
                    var v30 bool = true
                    fmt.Println(fi24, v30)
                    _ = v30
                }
                v23 -= -1
            }
            fmt.Println(v23, fi24)
            fmt.Println(v23)
        }
        _ = fi24
        _ = v25
    }
    _ = v23
    return ((("" + "hello") + "a") + "\n\t")
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 164
    var v3 int8 = 127
    var v4 string = (("\n\t" + (("hello" + ("\n\t" + "hello")) + ("" + "\n\t"))) + (("foo bar" + (("hello" + ("hello" + "a")) + "a")) + (("hello" + "\n\t") + ("" + "a"))))
    var v5 int8 = (v3 + (-(v3)))
    var v6 bool = true
    fmt.Println(v2, v5)
    var v31 uint = 730
    fmt.Println(v31, v3)
    var v32 bool = (-139 != 625)
    var v33 string = ("" + v4)
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            v5 = genFn0(v6)
            _ = fi35
        }
        v36 := GenS3{F0: 42}
        _ = fi34
        _ = v36
    }
    var v37 int16 = 1
    switch 0 {
    case 42:
        v3 += (v5 * v3)
    }
    delete(v1, "a")
    var v38 uint8 = 127
    v33 = ("a" + (v4 + v4))
    var v39 bool = v32
    var v40 bool = ((v37 + v37) > v37)
    var v41 int16 = 1
    v42 := GenS3{F0: 610}
    v2 -= (v2 ^ 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
