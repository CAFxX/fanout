package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int; F1 int8; F2 uint64 }
func genFn0() string {
    var v7 uint16 = 32767
    v8 := GenS1{F0: 65535, F1: 42}
    v7 = v7
    var v9 uint64 = 1
    if true {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int8 = 7
            var v12 string = ""
            fmt.Println(v9, v9)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return "a"
}
func genFn1() uint16 {
    if false {
        var v13 int32 = 2147483647
        fmt.Println(v13)
        _ = v13
    }
    if (!("a" != ("a" + "foo bar"))) {
        v14 := []int{}
        if false {
            var v15 bool = (false && (2147483647 > 111))
            for fi16 := 0; fi16 < 16; fi16++ {
                if (0.1 == 0.1) {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v15 = (0 == -1)
                        switch fi16 {
                        case -243:
                            v15 = v15
                        case -917:
                            v18 := GenS1{F0: 4294967295, F1: 4294967295}
                            _ = v18
                        case 4294967295:
                            var v19 int64 = 9223372036854775807
                            _ = v19
                        }
                        var v20 uint32 = 2147483648
                        _ = fi17
                        _ = v20
                    }
                } else {
                    v21 := GenS1{F0: 2, F1: 65535}
                    v15 = ((732 > 18446744073709551615) || (!v15))
                    var v22 bool = (!v15)
                    _ = v21
                    _ = v22
                }
                _ = fi16
            }
            var v23 uint = 560
            _ = v15
            _ = v23
        } else {
            var v24 uint8 = 13
            switch "hello" {
            case "":
                if (42 <= 18446744073709551615) {
                    v24 -= (v24 % (((v24 / ((v24) | 1))) | 1))
                    v14 = append(v14, 9223372036854775807)
                }
            default:
                if (float64(((float32((-2.25 / 1.5)) + (1e10 - 1.5)) - (1.5 * 1.5))) <= (0.1 * ((0.0 + 0.1) / 1.5))) {
                    v24 -= (v24 & 93)
                    if ((v24 == 42) || ((-2.25 == -2.25) && (false && ((1 > 0) && true)))) {
                        v24 = (^(v24))
                    } else {
                        v25 := GenS2{F0: 1}
                        _ = v25
                    }
                }
            }
            _ = v24
        }
        _ = v14
    }
    switch ("foo bar" + ("" + "\n\t")) {
    case "foo bar":
        var v26 int16 = 1000
        _ = v26
    case "":
        var v27 string = "foo bar"
        _ = v27
    }
    var v28 uint16 = 65535
    switch 617 {
    case -169:
        var v29 uint32 = 200
        _ = v29
    case 32767:
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := []int{fi30, (fi30 | fi30), fi30}
            v32 := GenS3{F0: 2, F1: 822}
            _ = fi30
            _ = v31
            _ = v32
        }
    case -418:
        v28 += (v28 | (v28 >> ((2) & 63)))
    }
    _ = v28
    return v28
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (((-2.25 * -2.25) * 0.1) + 1.5)
    var v3 int64 = -1
    var v4 int64 = (v3 * v3)
    v5 := []string{"", ("\n\t" + "a"), ("hello" + "")}
    v6 := make(map[int]string)
    var v33 uint8 = 80
    v34 := GenS2{F0: 2147483648}
    switch 794 {
    case 0:
        switch (("\n\t" + "\n\t") + "hello") {
        case "hello":
            var v35 int8 = -1
            _ = v35
        }
    case -1:
        var v36 uint64 = 1
        _ = v36
    }
    fmt.Println(v3, v33, v3)
    var v37 int = -614
    var v38 float32 = 0.1
    v39 := GenS3{F0: 42, F1: 299}
    var v40 int8 = 42
    var v41 bool = true
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS4{F0: 368, F1: v40, F2: 726}
        _ = fi42
        _ = v43
    }
    fmt.Println(v4, v33)
    _ = len(v1)
    fmt.Println(v38, v2, v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
