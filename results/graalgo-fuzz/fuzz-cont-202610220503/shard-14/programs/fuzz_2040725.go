package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 bool = (true && false)
    var v2 uint32 = 2147483648
    var v3 uint32 = 0
    var v4 int32 = -1
    v5 := make(map[uint32]int)
    var v6 bool = (v1 || true)
    for fi7 := 0; fi7 < 20; fi7++ {
        switch -1 {
        case -753:
            if v6 {
                for fi8 := 0; fi8 < 3; fi8++ {
                    if (1 >= 1) {
                        fmt.Println(v4, v4, fi8)
                        fmt.Println(v1)
                    }
                    switch "a" {
                    case "hello":
                        v1 = v6
                    default:
                        v9 := GenS1{F0: "\n\t"}
                        _ = v9
                    }
                    _ = fi8
                }
            } else {
                if (("a" + "\n\t") > "a") {
                    v6 = false
                    v1 = (int8(v4) == -1)
                    v10 := GenS1{F0: "\n\t"}
                    _ = v10
                } else {
                    var v11 uint64 = 606
                    fmt.Println(v6, v2)
                    _ = v11
                }
                v1 = ((("foo bar" + "\n\t") + "\n\t") < "\n\t")
            }
        case 127:
            for fi12 := 0; fi12 < 19; fi12++ {
                if v1 {
                    fmt.Println(v2, v1)
                    switch (fi7 % ((fi12) | 1)) {
                    case -312:
                        v3 -= (v2 >> ((255) & 63))
                    case -2147483648:
                        fmt.Println(v3, v6, fi12)
                    }
                    fmt.Println(v4, v6, v1)
                }
                _ = fi12
            }
        case -188:
            if (!true) {
                v4 += v4
                _ = len(v5)
                for fi13 := 0; fi13 < 19; fi13++ {
                    v1 = false
                    _ = fi13
                }
            }
        default:
            v14 := v5[v2]
            _ = v14
        }
        v15 := GenS1{F0: ("\n\t" + "hello")}
        _ = fi7
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS2{F0: 1, F1: v1, F2: 0}
        var v18 int = 2147483647
        _ = fi16
        _ = v17
        _ = v18
    }
    if v1 {
        switch 0 {
        case 2147483647:
            var v19 bool = (42 == 1)
            _ = v19
        }
        if ((732 == 1) && (32767 == 32767)) {
            if v6 {
                v20 := v5[(v3 * v2)]
                _ = v20
            }
            var v21 uint = 894
            _ = v21
        }
    }
    var v22 string = ("a" + "\n\t")
    v23 := GenS1{F0: v22}
    v24 := GenS2{F0: 1, F1: (((0.1 / -2.25) * (1e10 - 3.14159)) == (3.14159 - (0.1 + 1.5))), F2: -1}
    var v25 string = v22
    v22 = "foo bar"
    if (145 != 2) {
        if (127 >= 0) {
            if v1 {
                v26 := GenS3{F0: 1}
                for fi27 := 0; fi27 < 3; fi27++ {
                    v22 = ""
                    switch v25 {
                    case "foo bar":
                        v28 := GenS3{F0: 79}
                        _ = v28
                    case "":
                        fmt.Println(v2)
                    default:
                        if ((fi27 + (fi27 % ((fi27) | 1))) != fi27) {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                var v30 int32 = v4
                                var v31 uint64 = 1
                                v32 := GenS2{F0: -9223372036854775808, F1: v1, F2: -345}
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS2{F0: -769, F1: false, F2: 0}
                            _ = v33
                        }
                    }
                    fmt.Println(v22, v4, fi27)
                    _ = fi27
                }
                v22 = ("a" + ("a" + v22))
                _ = v26
            }
            switch 32767 {
            case 32768:
                var v34 int64 = 9223372036854775807
                _ = v34
            }
            fmt.Println(v4, v2)
        } else {
            for fi35 := 0; fi35 < 17; fi35++ {
                if false {
                    v22 = v25
                    var v36 int = (fi35 % ((fi35) | 1))
                    _ = v36
                }
                v4 = (v4 + (v4 << ((438) & 63)))
                v4 = v4
                _ = fi35
            }
            fmt.Println(v4, v6, v6)
            v4 -= (v4 ^ (v4 + (^(v4))))
        }
    } else {
        v1 = v1
    }
    v37 := []int{-675, 9223372036854775807, 4294967295, 128}
    v22 = ((("" + v22) + "hello") + v25)
    switch -827 {
    case 707:
        fmt.Println(v2, v22)
    case 32768:
        if v1 {
            var v38 bool = ((v4 % ((v4) | 1)) < (v4 << ((18446744073709551615) & 63)))
            _ = v38
        }
    default:
        fmt.Println(v6)
    }
    if v1 {
        switch v25 {
        case "hello":
            v39 := GenS4{F0: 0, F1: (v25 != ((v25 + v22) + ("\n\t" + v25))), F2: "a"}
            _ = v39
        case "\n\t":
            var v40 int64 = -9223372036854775808
            _ = v40
        default:
            if true {
                var v41 int8 = -1
                _ = v41
            }
        }
        var v42 int16 = -1
        fmt.Println(v3, v4)
        _ = v42
    }
    switch "hello" {
    case "foo bar":
        v2 += (v2 ^ v2)
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
