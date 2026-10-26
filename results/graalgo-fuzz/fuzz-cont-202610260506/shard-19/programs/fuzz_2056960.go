package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 bool }
func genFn0() int8 {
    if (127 < 66) {
        for fi4 := 0; fi4 < 18; fi4++ {
            var v5 uint8 = 128
            _ = fi4
            _ = v5
        }
        var v6 int64 = 1
        _ = v6
    } else {
        var v7 uint = 255
        var v8 int16 = -32768
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = ("hello" + "hello")
            var v11 int = fi9
            for fi12 := 0; fi12 < 19; fi12++ {
                if (-1 > 2147483647) {
                    v13 := []int{(v11 << ((v7) & 63)), fi9, (-(v11))}
                    if (!true) {
                        var v14 uint8 = 128
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        fmt.Println(fi12, fi9)
                        _ = v14
                    }
                    if (false && false) {
                        var v15 int32 = 521
                        var v16 int64 = 9223372036854775807
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                } else {
                    v17 := []int{fi12, 128, (fi9 - fi12)}
                    _ = len(v17)
                    v10 = "hello"
                    _ = v17
                }
                var v18 int = fi12
                _ = fi12
                _ = v18
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "a":
        var v19 int32 = 2147483647
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    var v21 uint8 = 142
    var v22 float32 = 1e10
    _ = v20
    _ = v21
    _ = v22
    return 42
}
func genFn1() int32 {
    for fi23 := 0; fi23 < 4; fi23++ {
        if true {
            var v24 bool = true
            v24 = v24
            _ = v24
        } else {
            for fi25 := 0; fi25 < 19; fi25++ {
                var v26 uint64 = 4294967295
                _ = fi25
                _ = v26
            }
        }
        _ = fi23
    }
    var v27 float32 = float32(3.14159)
    v27 = v27
    v28 := []int{32767}
    _ = v27
    _ = v28
    return -824
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int32 = -1
    if true {
        v3 = (v3 << ((0) & 63))
    } else {
        if (float32(0.0) > (-2.25 - (3.14159 / 1.5))) {
            v3 += -1
            v29 := v1[("hello" + ((("hello" + ("hello" + (("a" + "") + ("foo bar" + ("hello" + "\n\t"))))) + "foo bar") + "foo bar"))]
            _ = v29
        } else {
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 float32 = 1e10
                _ = fi30
                _ = v31
            }
            if true {
                v3 = genFn1()
            } else {
                var v32 int32 = v3
                v33 := GenS1{F0: -897}
                switch "" {
                case "":
                    if (65535 != 770) {
                        var v34 uint16 = 427
                        _ = v34
                    }
                case "a":
                    if (false || false) {
                        var v35 uint = 18446744073709551615
                        _ = v35
                    }
                case "hello":
                    for fi36 := 0; fi36 < 6; fi36++ {
                        for fi37 := 0; fi37 < 18; fi37++ {
                            fmt.Println(v32)
                            v38 := []int{2147483647, (fi36 & -9223372036854775808), (fi37 - (fi36 - (fi36 | fi37))), (^(fi36))}
                            _ = fi37
                            _ = v38
                        }
                        v3 += (v32 & (v3 | (v3 - genFn1())))
                        var v39 float32 = 0.1
                        _ = fi36
                        _ = v39
                    }
                }
                _ = v32
                _ = v33
            }
            if false {
                for fi40 := 0; fi40 < 9; fi40++ {
                    _ = len(v1)
                    _ = fi40
                }
                var v41 uint64 = 2
                _ = v41
            } else {
                switch "a" {
                case "":
                    if (573 == 526) {
                        v2[("\n\t" + "")] = ("foo bar" + ("" + "hello"))
                        var v42 int16 = -32768
                        for fi43 := 0; fi43 < 2; fi43++ {
                            v44 := GenS2{F0: "a"}
                            _ = fi43
                            _ = v44
                        }
                        _ = v42
                    }
                default:
                    var v45 bool = (65535 != 4294967295)
                    _ = v45
                }
                fmt.Println(v3)
            }
        }
    }
    v46 := GenS2{F0: ("foo bar" + ("\n\t" + "\n\t"))}
    var v47 int32 = 2147483647
    fmt.Println(v47, v3)
    for fi48 := 0; fi48 < 9; fi48++ {
        fmt.Println(v47, fi48)
        _ = fi48
    }
    var v49 uint16 = 112
    var v50 uint = 255
    switch "foo bar" {
    case "foo bar":
        for fi51 := 0; fi51 < 20; fi51++ {
            delete(v1, ((("\n\t" + "hello") + ("\n\t" + "\n\t")) + ""))
            _ = fi51
        }
    case "\n\t":
        switch ("hello" + "\n\t") {
        case "a":
            v3 = (v3 & 474)
        case "foo bar":
            v52 := []int{}
            _ = v52
        }
    case "hello":
        v49 = (^(v49))
    default:
        v3 = (v47 / (((^(v47))) | 1))
    }
    var v53 int64 = 42
    switch "a" {
    case "a":
        for fi54 := 0; fi54 < 2; fi54++ {
            v55 := GenS3{F0: v50, F1: false}
            var v56 uint = 255
            switch (-(fi54)) {
            case 4294967295:
                v57 := []int{fi54}
                _ = v57
            case 836:
                for fi58 := 0; fi58 < 13; fi58++ {
                    v53 = v53
                    v59 := []int{(fi54 | fi54), fi58, fi58, fi54}
                    _ = fi58
                    _ = v59
                }
            }
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    fmt.Println(v53)
    var v60 int32 = int32(v50)
    for fi61 := 0; fi61 < 9; fi61++ {
        v62 := []int{}
        _ = len(v2)
        _ = fi61
        _ = v62
    }
    if false {
        v60 = (v47 >> ((v50) & 63))
    }
    v63 := GenS2{F0: ("" + ("" + ("hello" + (("" + ("hello" + "")) + "a"))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v49
    _ = v50
    _ = v53
    _ = v60
    _ = v63
}
