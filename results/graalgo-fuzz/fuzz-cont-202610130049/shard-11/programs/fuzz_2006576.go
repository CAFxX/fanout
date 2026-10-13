package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint; F2 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 string; F1 uint32; F2 uint8 }
func genFn0(p0 uint64, p1 bool) int8 {
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            if true {
                fmt.Println(p0)
            }
            _ = fi6
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 int, p1 int16, p2 int8) uint32 {
    v7 := []int{4294967295, p0, p0}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p1 -= p1
    if false {
        var v8 string = "foo bar"
        _ = v8
    } else {
        var v9 int64 = 42
        p1 = p1
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int8 = 127
    v3 := make(map[int]string)
    if true {
        if (4294967295 < 1) {
            var v10 uint64 = 443
            _ = v10
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 uint16 = 2
            v13 := []int{(^(fi11)), fi11}
            v1 = "foo bar"
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    switch v1 {
    case "":
        switch v1 {
        case "foo bar":
            var v14 float32 = 1e10
            _ = v14
        default:
            if ((true && true) && true) {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v2 -= v2
                    if true {
                        _ = len(v3)
                    } else {
                        v16 := v3[(fi15 ^ (fi15 << ((65535) & 63)))]
                        switch v1 {
                        case "hello":
                            v17 := GenS1{F0: -93}
                            _ = v17
                        case "foo bar":
                            if false {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    v19 := GenS1{F0: 42}
                                    _ = fi18
                                    _ = v19
                                }
                                if (-1 == -1) {
                                    v20 := GenS1{F0: -1}
                                    _ = v20
                                }
                                var v21 int64 = -9223372036854775808
                                _ = v21
                            }
                        case "":
                            switch (-(fi15)) {
                            case 127:
                                v22 := GenS2{F0: 1, F1: 42}
                                _ = v22
                            case 0:
                                v23 := []int{(fi15 >> ((65535) & 63)), fi15}
                                _ = v23
                            case 42:
                                v2 -= (-(v2))
                            }
                        }
                        v24 := []int{9223372036854775807, -628, (-(fi15)), 127}
                        _ = v16
                        _ = v24
                    }
                    _ = fi15
                }
                v25 := []int{-145, -1}
                if false {
                    v26 := GenS3{F0: 364}
                    v1 = ((v1 + v1) + ((v1 + "hello") + v1))
                    _ = v26
                }
                _ = v25
            }
        }
    case "a":
        var v27 int = -423
        _ = v27
    case "hello":
        if (false && true) {
            v28 := GenS3{F0: 1}
            switch -799 {
            case 9223372036854775807:
                var v29 int = 32768
                _ = v29
            }
            _ = v28
        } else {
            v2 = (-(v2))
            var v30 int8 = (v2 >> ((18446744073709551615) & 63))
            _ = v30
        }
    }
    var v31 uint64 = 4294967295
    fmt.Println(v31, v31, v31)
    if (v2 < (v2 + 42)) {
        v2 = genFn0(4294967295, (v1 <= v1))
        var v32 float32 = (3.14159 * 1e10)
        v2 = -80
        _ = v32
    } else {
        v33 := GenS1{F0: -822}
        v34 := GenS4{F0: -32768, F1: 2, F2: -148}
        _ = v33
        _ = v34
    }
    var v35 float32 = (float32((1.5 / (1e10 * 3.14159))) * 0.1)
    var v36 int16 = 42
    v37 := GenS5{F0: -2147483648, F1: (v31 - (^(v31)))}
    switch "" {
    case "foo bar":
        v38 := []int{-1, 128, -573}
        _ = v38
    case "a":
        if false {
            switch 9223372036854775807 {
            case 127:
                fmt.Println(v2, v36)
            case -2147483648:
                _ = len(v3)
            case 2147483647:
                for fi39 := 0; fi39 < 9; fi39++ {
                    v1 = "\n\t"
                    var v40 uint32 = 492
                    v36 = 42
                    _ = fi39
                    _ = v40
                }
            }
            switch ((v1 + v1) + v1) {
            case "foo bar":
                switch 1 {
                case 42:
                    v41 := GenS4{F0: 32767, F1: 18446744073709551615, F2: 9223372036854775807}
                    _ = v41
                case 4294967295:
                    v42 := GenS2{F0: 42, F1: v31}
                    _ = v42
                case 32768:
                    for fi43 := 0; fi43 < 19; fi43++ {
                        v36 = v36
                        _ = fi43
                    }
                default:
                    var v44 uint16 = 1
                    _ = v44
                }
            }
            fmt.Println(v36, v31, v1)
        }
    }
    switch v1 {
    case "":
        for fi45 := 0; fi45 < 19; fi45++ {
            if false {
                v1 = (v1 + v1)
                var v46 bool = ((9223372036854775807 != -617) && true)
                _ = v46
            }
            _ = fi45
        }
    case "\n\t":
        if (-237 >= 0) {
            v1 = "a"
            v1 = v1
        } else {
            v47 := GenS3{F0: 42}
            for fi48 := 0; fi48 < 8; fi48++ {
                delete(v3, fi48)
                var v49 uint32 = 2147483647
                if true {
                    switch 128 {
                    case 9223372036854775807:
                        v49 += v49
                    case 134:
                        v50 := GenS6{F0: "hello", F1: (v49 * v49), F2: 1}
                        _ = v50
                    default:
                        var v51 int32 = -2147483648
                        _ = v51
                    }
                }
                _ = fi48
                _ = v49
            }
            _ = v47
        }
    }
    v52 := []int{467, -2147483648, 32768}
    v1 = v1
    var v53 int16 = v36
    delete(v3, 32768)
    var v54 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v52
    _ = v53
    _ = v54
}
