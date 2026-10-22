package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 string }
type GenS5 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS6 struct { F0 int8; F1 bool; F2 int64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            fmt.Println(fi7)
            switch fi7 {
            case -776:
                fmt.Println(fi7)
            case 128:
                v8 := []int{45, fi7, (fi7 * fi7)}
                _ = v8
            }
        } else {
            v9 := GenS1{F0: 1, F1: -1}
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{128}
    if true {
        var v11 uint64 = 926
        v11 = 750
        for fi12 := 0; fi12 < 15; fi12++ {
            v11 -= 390
            v11 = v11
            _ = fi12
        }
        _ = v11
    }
    if false {
        if false {
            var v13 int16 = 42
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int = (fi14 << ((255) & 63))
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS2{F0: "a", F1: 723, F2: (v15 / ((-1) | 1))}
                    if len(v10) > 0 {
                        v10[0] = (fi16 / ((fi16) | 1))
                    }
                    var v18 int64 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                fmt.Println(v15)
                _ = fi14
                _ = v15
            }
            v13 -= v13
            _ = v13
        }
        var v19 uint16 = 32767
        var v20 string = "foo bar"
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: "\n\t", F1: 42, F2: (^(fi21))}
        var v23 uint64 = 2
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = v10
    return 100
}
func genFn1(p0 int, p1 uint64, p2 uint32) int {
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(p2, p2)
        for fi25 := 0; fi25 < 10; fi25++ {
            switch (^(fi25)) {
            case 601:
                fmt.Println(p2, fi24, fi25)
            case 42:
                v26 := []int{1, fi25, (^(fi24)), (^(fi24))}
                if (-9223372036854775808 < -9223372036854775808) {
                    p2 = (p2 ^ (^(p2)))
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    v27 := GenS3{F0: 65535, F1: 32767}
                    _ = v27
                } else {
                    if false {
                        switch -802 {
                        case 42:
                            var v28 int32 = 960
                            _ = v28
                        }
                        switch "" {
                        case "foo bar":
                            var v29 bool = false
                            _ = v29
                        }
                    } else {
                        if false {
                            fmt.Println(p2, p1)
                            if (false && true) {
                                v30 := GenS4{F0: p2, F1: p2, F2: "a"}
                                _ = v30
                            }
                            var v31 int = fi24
                            v32 := GenS5{F0: 128, F1: 65535, F2: 187}
                            _ = v31
                            _ = v32
                        }
                        p1 = 42
                        var v33 uint64 = (^(p1))
                        _ = v33
                    }
                    v26 = append(v26, (p0 / ((-9223372036854775808) | 1)))
                    v34 := GenS1{F0: -540, F1: 0}
                    _ = v34
                }
                _ = v26
            case 128:
                var v35 float64 = 0.1
                _ = v35
            }
            _ = fi25
        }
        _ = fi24
    }
    p0 -= (-(p0))
    _ = p0
    _ = p1
    _ = p2
    return (p0 - -9223372036854775808)
}

func main() {
    var v1 uint16 = 32767
    var v2 string = ""
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v6 uint8 = 2
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 uint = 65535
        v38 := GenS5{F0: (v6 << ((v37) & 63)), F1: (v1 << ((4294967295) & 63)), F2: 42}
        v39 := []int{32767, fi36}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v6 += (v6 - (v6 - 127))
    v40 := GenS5{F0: (v6 ^ uint8(v1)), F1: (v1 + v1), F2: 0}
    var v41 uint32 = 0
    var v42 uint = 1
    v6 += 1
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 int8 = 0
        _ = fi43
        _ = v44
    }
    v45 := GenS6{F0: 1, F1: false, F2: 261}
    v46 := GenS5{F0: 96, F1: 953, F2: 42}
    v2 = v2
    var v47 uint32 = v41
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
