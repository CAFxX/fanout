package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint8; F2 string }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 int32) int64 {
    fmt.Println(p1)
    switch 32768 {
    case 2147483647:
        var v7 uint32 = 0
        _ = v7
    case 1:
        if p0 {
            p0 = (2 <= 2)
            if p0 {
                var v8 uint16 = 65535
                if (1 < 32767) {
                    var v9 uint8 = uint8(v8)
                    v10 := []int{}
                    v11 := GenS1{F0: -1, F1: -2147483648}
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v12 := []int{int(v8), 0}
                    if p0 {
                        p0 = (363 > 2147483647)
                        switch 2147483647 {
                        case 127:
                            v13 := GenS1{F0: 127, F1: (p1 >> ((2) & 63))}
                            _ = v13
                        }
                    } else {
                        switch (("hello" + ("\n\t" + ("\n\t" + "hello"))) + "hello") {
                        case "foo bar":
                            p1 = (p1 ^ p1)
                        case "":
                            var v14 float32 = 0.1
                            _ = v14
                        }
                    }
                    _ = v12
                }
                _ = v8
            }
        }
    case -2147483648:
        v15 := GenS2{F0: 1}
        _ = v15
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 int32, p1 bool) uint {
    var v16 int = 127
    if p1 {
        switch "foo bar" {
        case "a":
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 19; fi18++ {
                    fmt.Println(p1)
                    _ = fi18
                }
                fmt.Println(v16, p0, v16)
                p0 = p0
                _ = fi17
            }
        case "":
            for fi19 := 0; fi19 < 18; fi19++ {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        v20 := []int{fi19}
                        _ = v20
                    case "":
                        fmt.Println(p1, v16, fi19)
                    case "hello":
                        p0 = p0
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v16 -= 849
                        v22 := GenS3{F0: 42, F1: 128}
                        var v23 bool = (998 != 2)
                        p1 = true
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                } else {
                    if ((p1 || (false && p1)) || (v16 < v16)) {
                        v24 := GenS4{F0: true}
                        _ = v24
                    }
                }
                _ = fi19
            }
        }
        var v25 int64 = 0
        _ = v25
    } else {
        v26 := GenS3{F0: 1000, F1: 1}
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v16
    return 540
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 13
    var v3 bool = true
    var v4 bool = v3
    v5 := make(map[int]bool)
    var v6 uint64 = 42
    var v27 int = int(v2)
    var v28 uint32 = 2147483648
    v29 := GenS5{F0: (v6 & uint64(v2)), F1: v2, F2: "\n\t"}
    if (0 <= 669) {
        v4 = (!(!(42 < 100)))
        var v30 float64 = 1e10
        v3 = (v4 || ("" == ("" + "\n\t")))
        _ = v30
    }
    var v31 int8 = -1
    switch "hello" {
    case "hello":
        v6 = 2
    case "foo bar":
        var v32 int16 = (v1 | v1)
        _ = v32
    case "\n\t":
        if v4 {
            v33 := GenS5{F0: 28, F1: uint8(v28), F2: "hello"}
            _ = v33
        }
    }
    v28 = v28
    var v34 float32 = 3.14159
    if ((v2 >> ((255) & 63)) != (v2 | v2)) {
        v35 := GenS1{F0: (v31 / (((-(v31))) | 1)), F1: 42}
        _ = len(v5)
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v3 = (!v3)
        for fi37 := 0; fi37 < 4; fi37++ {
            switch "" {
            case "foo bar":
                v38 := GenS6{F0: (^(v2)), F1: (v6 / (((v6 / (((v6 >> ((2) & 63))) | 1))) | 1))}
                _ = v38
            case "":
                v39 := v5[v27]
                _ = v39
            }
            _ = fi37
        }
        if (!true) {
            v40 := []int{v27, 128, fi36}
            v41 := GenS7{F0: ("a" + "\n\t"), F1: 2}
            if true {
                var v42 int = 2147483647
                if (!(!v4)) {
                    switch "foo bar" {
                    case "a":
                        for fi43 := 0; fi43 < 20; fi43++ {
                            var v44 uint8 = uint8(v28)
                            _ = len(v5)
                            if (v44 < v44) {
                                if (1 != 65535) {
                                    if (42 > v44) {
                                        var v45 bool = (v6 >= v6)
                                        _ = v45
                                    }
                                    fmt.Println(v3)
                                } else {
                                    var v46 bool = true
                                    _ = v46
                                }
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        v47 := GenS3{F0: (v1 >> ((43) & 63)), F1: (^(v2))}
                        _ = v47
                    }
                    var v48 uint8 = v2
                    v2 -= v2
                    _ = v48
                }
                v49 := GenS8{F0: v31, F1: 0}
                _ = v42
                _ = v49
            }
            _ = v40
            _ = v41
        } else {
            fmt.Println(v6, v1, v3)
        }
        _ = fi36
    }
    v6 -= (v6 + v6)
    for fi50 := 0; fi50 < 6; fi50++ {
        var v51 int16 = -32768
        var v52 uint16 = 0
        for fi53 := 0; fi53 < 18; fi53++ {
            v34 = ((v34 / float32(3.14159)) * (v34 * v34))
            var v54 float64 = -2.25
            _ = fi53
            _ = v54
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
