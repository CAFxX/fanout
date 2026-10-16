package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint; F2 string }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 string = ("a" + ("hello" + ""))
    var v4 int = 127
    var v5 uint32 = 567
    for fi6 := 0; fi6 < 17; fi6++ {
        delete(v1, (v4 | fi6))
        _ = fi6
    }
    var v7 int8 = 68
    if (-126 != v7) {
        if (!(0 != 118)) {
            v8 := []int{(v4 | (-(v4))), v4, 1, 32768}
            _ = v8
        } else {
            if false {
                var v9 float64 = (1e10 - 3.14159)
                for fi10 := 0; fi10 < 20; fi10++ {
                    v9 += v9
                    for fi11 := 0; fi11 < 15; fi11++ {
                        v12 := GenS1{F0: 1}
                        if false {
                            var v13 int = fi11
                            v4 -= 9223372036854775807
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                    if true {
                        switch v3 {
                        case "hello":
                            for fi14 := 0; fi14 < 14; fi14++ {
                                v1[-2147483648] = "hello"
                                v5 += (v5 + v5)
                                var v15 bool = (float32((v9 + 0.0)) != 0.0)
                                _ = fi14
                                _ = v15
                            }
                        case "a":
                            for fi16 := 0; fi16 < 13; fi16++ {
                                var v17 float32 = 1.5
                                v18 := GenS2{F0: (v5 % ((v5) | 1))}
                                if (1 < 42) {
                                    v7 -= v7
                                    v19 := GenS1{F0: 79}
                                    if (!true) {
                                        v20 := []int{fi16, 313, 684}
                                        _ = v20
                                    } else {
                                        v21 := GenS2{F0: v5}
                                        _ = v21
                                    }
                                    _ = v19
                                } else {
                                    v22 := []int{}
                                    var v23 int16 = 0
                                    var v24 uint16 = 1
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        case "":
                            if (v7 <= v7) {
                                fmt.Println(fi10)
                                delete(v1, v4)
                            } else {
                                fmt.Println(v5)
                                fmt.Println(v3, v3, fi10)
                            }
                        default:
                            v4 = -1
                        }
                        v25 := GenS3{F0: 65535, F1: 65535, F2: ((v3 + v3) + "")}
                        _ = v25
                    }
                    _ = fi10
                }
                fmt.Println(v4, v3, v7)
                _ = v9
            } else {
                fmt.Println(v3, v7, v7)
            }
            if (!true) {
                var v26 uint64 = 4294967295
                v4 -= int(v7)
                var v27 int32 = 42
                _ = v26
                _ = v27
            } else {
                var v28 int32 = -394
                _ = v28
            }
            switch (v3 + "") {
            case "\n\t":
                v1[v4] = "\n\t"
            case "":
                var v29 string = (v3 + (("hello" + (((v3 + v3) + v3) + v3)) + v3))
                _ = v29
            default:
                v30 := GenS4{F0: -972}
                _ = v30
            }
        }
        switch v3 {
        case "foo bar":
            var v31 bool = false
            _ = v31
        case "\n\t":
            v32 := []int{}
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 3; fi33++ {
                switch "" {
                case "a":
                    fmt.Println(fi33, v4)
                }
                _ = fi33
            }
        default:
            var v34 int8 = v7
            _ = v34
        }
    }
    v35 := []int{}
    for fi36 := 0; fi36 < 7; fi36++ {
        if true {
            for fi37 := 0; fi37 < 16; fi37++ {
                fmt.Println(v7, v5, v3)
                _ = fi37
            }
            var v38 bool = (42 > 902)
            switch fi36 {
            case 2147483647:
                v4 = (fi36 >> ((255) & 63))
            case -1:
                v4 -= v4
            case 264:
                var v39 int16 = -1
                _ = v39
            }
            _ = v38
        } else {
            v40 := GenS2{F0: v5}
            if false {
                v35 = append(v35, (-(fi36)))
                v5 += 485
            } else {
                var v41 uint16 = 2
                v3 = (("foo bar" + v3) + ("hello" + ""))
                _ = v41
            }
            _ = v40
        }
        v42 := GenS3{F0: 65535, F1: 4294967295, F2: "a"}
        _ = fi36
        _ = v42
    }
    fmt.Println(v3, v5, v3)
    v43 := v1[(v4 | (v4 % ((v4) | 1)))]
    for fi44 := 0; fi44 < 4; fi44++ {
        v43 = ""
        var v45 string = (v3 + ((v43 + ((v3 + "a") + (("hello" + v3) + v43))) + ("" + (("a" + (v43 + "")) + v43))))
        _ = fi44
        _ = v45
    }
    var v46 int8 = (^(v7))
    for fi47 := 0; fi47 < 14; fi47++ {
        var v48 bool = true
        for fi49 := 0; fi49 < 6; fi49++ {
            v50 := GenS2{F0: 2}
            if v48 {
                if false {
                    v46 = (v7 >> ((255) & 63))
                }
                var v51 int8 = v7
                _ = v51
            }
            v43 = ("foo bar" + "\n\t")
            _ = fi49
            _ = v50
        }
        for fi52 := 0; fi52 < 19; fi52++ {
            v48 = false
            fmt.Println(v3)
            _ = fi52
        }
        _ = fi47
        _ = v48
    }
    var v53 uint32 = v5
    if false {
        for fi54 := 0; fi54 < 5; fi54++ {
            var v55 int8 = (v46 - 127)
            var v56 int = fi54
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v53
    _ = v7
}
