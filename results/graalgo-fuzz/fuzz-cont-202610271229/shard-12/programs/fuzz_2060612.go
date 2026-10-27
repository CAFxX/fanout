package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint32) uint8 {
    var v7 uint32 = (p1 ^ p1)
    p1 += (v7 ^ (p1 + p1))
    fmt.Println(p1, v7, p1)
    if (5 <= 93) {
        var v8 int8 = -128
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: "\n\t"}
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int8 = 0
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v13 bool = false
        _ = v13
        _ = v8
    } else {
        p1 += 2147483647
        v14 := []int{2147483647, 1, 2147483647, 1}
        for fi15 := 0; fi15 < 17; fi15++ {
            _ = len(v14)
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: ""}
                _ = len(v14)
                if true {
                    var v18 bool = (203 >= 2)
                    if (163 < 1) {
                        v19 := GenS1{F0: ("\n\t" + "hello")}
                        v7 -= 2147483647
                        fmt.Println(v18, v18)
                        _ = v19
                    } else {
                        if len(v14) > 0 {
                            v14[0] = (fi16 % (((fi16 * (fi16 + fi16))) | 1))
                        }
                        var v20 uint8 = 1
                        _ = v20
                    }
                    _ = v18
                } else {
                    v21 := GenS2{F0: 18446744073709551615, F1: "", F2: -1}
                    _ = v21
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        _ = v14
    }
    switch -2147483648 {
    case 9223372036854775807:
        p0 -= (-(p0))
    default:
        if true {
            v22 := GenS2{F0: 65535, F1: "hello", F2: 0}
            for fi23 := 0; fi23 < 15; fi23++ {
                if true {
                    v24 := GenS3{F0: -32768, F1: -26}
                    _ = v24
                }
                _ = fi23
            }
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 uint16) int64 {
    fmt.Println(p0, p0)
    var v25 string = "a"
    var v26 float32 = (1e10 / 1e10)
    p0 = 686
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 0
}

func main() {
    v1 := []string{("foo bar" + "a"), ""}
    v2 := make(map[int]uint32)
    var v3 uint = 0
    v4 := make(map[int]uint32)
    var v5 uint8 = 79
    v6 := make(map[string]int)
    fmt.Println(v5, v3)
    var v27 int = 1
    v5 = (v5 >> ((v3) & 63))
    var v28 uint16 = 687
    v29 := GenS4{F0: 127, F1: 1}
    var v30 string = "hello"
    var v31 uint32 = 2147483647
    v32 := GenS2{F0: v3, F1: v30, F2: int32(v31)}
    if false {
        delete(v4, -1)
        if (true && true) {
            v33 := v4[v27]
            if true {
                switch v27 {
                case 365:
                    var v34 uint = v3
                    _ = v34
                }
                v35 := GenS2{F0: 65535, F1: v30, F2: 343}
                v3 = (v3 ^ (v3 + (v3 ^ 512)))
                _ = v35
            } else {
                var v36 string = v30
                _ = v36
            }
            _ = v33
        } else {
            for fi37 := 0; fi37 < 17; fi37++ {
                v1 = append(v1, v30)
                switch (fi37 % (((v27 | v27)) | 1)) {
                case 32767:
                    var v38 uint = (v3 >> ((255) & 63))
                    _ = v38
                case -1:
                    var v39 float32 = 3.14159
                    _ = v39
                }
                _ = fi37
            }
            if ((-98 != 100) && true) {
                var v40 uint = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v40
            } else {
                for fi41 := 0; fi41 < 8; fi41++ {
                    v28 = (v28 * v28)
                    if false {
                        v30 = v30
                    } else {
                        v42 := GenS5{F0: (^(v31))}
                        _ = v42
                    }
                    _ = fi41
                }
                var v43 uint16 = (^(v28))
                _ = v43
            }
        }
    } else {
        v28 = v28
        var v44 int32 = -420
        _ = v44
    }
    if true {
        for fi45 := 0; fi45 < 11; fi45++ {
            fmt.Println(v30)
            _ = fi45
        }
        v46 := GenS4{F0: 100, F1: 1}
        v30 = v30
        _ = v46
    }
    _ = len(v6)
    if len(v1) > 0 {
        v1[0] = (("foo bar" + "hello") + (v30 + ((v30 + ((v30 + ("foo bar" + "")) + v30)) + v30)))
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v5 = 127
    fmt.Println(v31, v30, v3)
    for fi47 := 0; fi47 < 5; fi47++ {
        v5 = 127
        _ = fi47
    }
    v28 = (v28 & 515)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
