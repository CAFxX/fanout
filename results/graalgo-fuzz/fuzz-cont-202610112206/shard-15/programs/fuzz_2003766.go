package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 int64, p1 string) int64 {
    v7 := []int{}
    fmt.Println(p1, p0)
    v7 = append(v7, 4294967295)
    v8 := GenS1{F0: p0, F1: 0}
    var v9 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p0 ^ 969)
}

func main() {
    var v1 int16 = -1
    var v2 int8 = -109
    var v3 float64 = (0.1 * 3.14159)
    v4 := make(map[int]int)
    var v5 float32 = (-2.25 / 3.14159)
    v6 := make(map[int]int)
    switch 9223372036854775807 {
    case 0:
        var v10 bool = (721 <= 0)
        _ = v10
    case 127:
        if true {
            if (!false) {
                var v11 uint64 = 833
                v3 += v3
                v12 := GenS1{F0: -9223372036854775808, F1: 815}
                _ = v11
                _ = v12
            }
            var v13 uint = 986
            v14 := GenS2{F0: 2147483648, F1: 2}
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 int32 = 2147483647
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 13; fi18++ {
                    switch (fi18 >> ((18446744073709551615) & 63)) {
                    case 127:
                        v19 := GenS1{F0: -788, F1: fi18}
                        _ = v19
                    default:
                        v20 := []int{}
                        _ = v20
                    }
                    _ = fi18
                }
                _ = fi15
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                v6[(fi21 * (fi21 & fi21))] = (fi21 >> ((950) & 63))
                _ = fi21
            }
        }
    case 9223372036854775807:
        if (false && true) {
            if (true && true) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    v6[(fi22 - (-(fi22)))] = (^(fi22))
                    switch 336 {
                    case -1:
                        v23 := []int{fi22, fi22}
                        _ = v23
                    case 762:
                        v24 := GenS2{F0: 42, F1: 449}
                        _ = v24
                    }
                    _ = fi22
                }
                v25 := GenS3{F0: 103, F1: ("a" + "hello"), F2: 42}
                var v26 uint = 373
                _ = v25
                _ = v26
            } else {
                v2 = (v2 << ((42) & 63))
                switch -9223372036854775808 {
                case 4294967295:
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 uint = 0
                        delete(v6, int(v2))
                        _ = fi27
                        _ = v28
                    }
                }
            }
            v29 := GenS1{F0: 42, F1: 32768}
            switch ("\n\t" + "hello") {
            case "foo bar":
                var v30 int16 = (-(v1))
                _ = v30
            case "\n\t":
                v1 = v1
            case "hello":
                if (42 > 1) {
                    v5 = 3.14159
                    v2 -= (v2 * (v2 >> ((317) & 63)))
                    v2 = 87
                } else {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        var v32 int16 = (^(v1))
                        var v33 bool = false
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(v3, v5, v5)
                }
            default:
                var v34 int64 = -9223372036854775808
                _ = v34
            }
            _ = v29
        } else {
            var v35 int64 = 42
            if false {
                var v36 uint8 = 173
                v37 := GenS4{F0: 4294967295, F1: 4294967295}
                fmt.Println(v3)
                _ = v36
                _ = v37
            }
            var v38 bool = true
            _ = v35
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        for fi40 := 0; fi40 < 18; fi40++ {
            v41 := GenS2{F0: 2147483648, F1: 0}
            var v42 bool = (42 >= 42)
            if true {
                switch 127 {
                case -883:
                    v43 := GenS3{F0: v2, F1: ("" + ""), F2: 425}
                    _ = v43
                case 32768:
                    switch fi40 {
                    case 127:
                        switch (fi39 << ((4294967295) & 63)) {
                        case -9223372036854775808:
                            v1 -= (v1 - v1)
                        case 127:
                            var v44 int64 = 0
                            _ = v44
                        default:
                            v5 = (v5 - 0.0)
                        }
                    case 9223372036854775807:
                        var v45 int64 = 0
                        _ = v45
                    }
                }
                var v46 uint = 0
                var v47 uint16 = 0
                _ = v46
                _ = v47
            } else {
                var v48 int64 = -591
                _ = v48
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    v1 = (v1 - (v1 % ((-839) | 1)))
    v49 := GenS1{F0: 4, F1: -2147483648}
    var v50 string = ""
    fmt.Println(v1, v2)
    v51 := v6[9223372036854775807]
    var v52 float64 = float64(v5)
    if false {
        fmt.Println(v2, v1, v3)
        v52 += v52
        var v53 uint32 = 164
        _ = v53
    } else {
        fmt.Println(v50)
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        if (776 >= 2) {
            v5 = (-2.25 * 0.0)
        }
        switch (((v50 + v50) + v50) + v50) {
        case "a":
            var v55 uint8 = 1
            _ = v55
        }
        _ = fi54
    }
    if (42 >= 18446744073709551615) {
        var v56 uint64 = 397
        var v57 string = v50
        var v58 int16 = int16(v51)
        _ = v56
        _ = v57
        _ = v58
    } else {
        var v59 float32 = v5
        v60 := GenS4{F0: 302, F1: 0}
        var v61 float32 = (3.14159 / -2.25)
        _ = v59
        _ = v60
        _ = v61
    }
    v50 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
