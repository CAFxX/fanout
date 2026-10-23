package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 string; F2 int }
type GenS4 struct { F0 string; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS9 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    switch 127 {
    case -693:
        for fi4 := 0; fi4 < 20; fi4++ {
            v5 := []int{}
            if ((false && (true || (1.5 > 0.0))) || false) {
                var v6 int = 2147483647
                if (p0 >= (p0 - 4294967295)) {
                    v7 := GenS1{F0: fi4, F1: 2}
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        fmt.Println(v6)
                        if len(v5) > 0 {
                            v5[0] = (fi8 & (-(v6)))
                        }
                        fmt.Println(fi8, v6)
                        _ = fi8
                    }
                    var v9 uint64 = 18446744073709551615
                    _ = v9
                }
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: (fi4 - (fi4 % ((v6) | 1))), F1: 751}
                    v12 := GenS2{F0: 363}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v6
            }
            _ = fi4
            _ = v5
        }
    case 9223372036854775807:
        p0 += p0
    case 127:
        if ((("a" + "\n\t") + "a") != "a") {
            p0 = (^(p0))
            v13 := GenS3{F0: 42, F1: ("foo bar" + "foo bar"), F2: 32768}
            _ = v13
        }
    }
    v14 := []int{0, 32768}
    var v15 int16 = 527
    _ = p0
    _ = v14
    _ = v15
    return 0
}
func genFn1(p0 bool, p1 uint64, p2 int32) bool {
    fmt.Println(p1)
    var v16 uint8 = 0
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: uint16(v16)}
        _ = fi17
        _ = v18
    }
    if p0 {
        var v19 uint32 = 2
        var v20 uint8 = (v16 % (((v16 * (^(v16)))) | 1))
        _ = v19
        _ = v20
    }
    if p0 {
        v21 := GenS1{F0: 159, F1: p1}
        v22 := GenS4{F0: "", F1: 1, F2: 149}
        _ = v21
        _ = v22
    } else {
        v23 := []int{}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (!((^(p2)) == p2))
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 0
    var v3 bool = false
    if true {
        v1 = (v1 ^ (v1 + (v1 | v1)))
        for fi24 := 0; fi24 < 15; fi24++ {
            if v3 {
                var v25 int64 = 9223372036854775807
                v3 = (-2.25 <= 0.1)
                var v26 uint32 = 2147483648
                _ = v25
                _ = v26
            } else {
                v2 = (v2 * (v2 | (v2 - v2)))
                v27 := []int{(fi24 ^ fi24)}
                _ = len(v27)
                _ = v27
            }
            v2 -= (^(v2))
            _ = fi24
        }
        v28 := []int{}
        _ = v28
    } else {
        var v29 string = ("" + "hello")
        _ = v29
    }
    switch "" {
    case "a":
        var v30 uint16 = 323
        _ = v30
    case "hello":
        v31 := []int{}
        _ = v31
    case "\n\t":
        if v3 {
            switch -808 {
            case -2147483648:
                v3 = (255 < 809)
            case 127:
                v32 := GenS5{F0: 255}
                _ = v32
            case 955:
                switch ("a" + "") {
                case "hello":
                    v1 = v1
                }
            }
        } else {
            v1 -= 255
        }
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        fmt.Println(v3)
        v34 := GenS2{F0: 439}
        _ = fi33
        _ = v34
    }
    v35 := []int{9223372036854775807}
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS6{F0: 42, F1: 549, F2: 460}
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := GenS1{F0: (fi36 | -162), F1: 1}
            if (v1 < (v1 | 127)) {
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
                if true {
                    v3 = v3
                    v40 := GenS6{F0: v1, F1: 65535, F2: uint64(v1)}
                    v35 = append(v35, (-(fi36)))
                    _ = v40
                } else {
                    v41 := GenS7{F0: 2}
                    if ((!(!v3)) && v3) {
                        fmt.Println(v1)
                        v3 = (((!((v2 / ((v2) | 1)) == 376)) || v3) && v3)
                        v42 := GenS1{F0: (-(fi38)), F1: v2}
                        _ = v42
                    } else {
                        var v43 int16 = -32768
                        _ = v43
                    }
                    var v44 float64 = float64(0.0)
                    _ = v41
                    _ = v44
                }
            }
            _ = fi38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    fmt.Println(v3)
    var v45 int8 = int8(v2)
    var v46 uint = 331
    v47 := GenS8{F0: 0, F1: 2, F2: (v2 >> (((^(v46))) & 63))}
    for fi48 := 0; fi48 < 16; fi48++ {
        fmt.Println(v2, v3, v2)
        switch (fi48 / ((int(v46)) | 1)) {
        case 920:
            v2 = (v2 + v2)
        }
        _ = fi48
    }
    v49 := GenS9{F0: ("" > "foo bar"), F1: -573}
    if v3 {
        for fi50 := 0; fi50 < 13; fi50++ {
            fmt.Println(v46, fi50)
            _ = fi50
        }
        v45 = v45
    }
    _ = len(v35)
    v46 = (v46 << ((uint(v2)) & 63))
    v1 = v1
    var v51 int8 = (-(v45))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v51
}
