package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 uint64) int {
    if (false && true) {
        for fi4 := 0; fi4 < 3; fi4++ {
            for fi5 := 0; fi5 < 14; fi5++ {
                if false {
                    var v6 int32 = 341
                    v6 = 42
                    _ = v6
                }
                var v7 bool = (false && (!false))
                fmt.Println(p0)
                _ = fi5
                _ = v7
            }
            _ = fi4
        }
        var v8 int = 127
        p0 = (p0 << ((4294967295) & 63))
        _ = v8
    } else {
        v9 := GenS1{F0: 4294967295, F1: 0, F2: int(p0)}
        if false {
            fmt.Println(p0)
        }
        _ = v9
    }
    switch ("hello" + "hello") {
    case "\n\t":
        var v10 int32 = -1
        _ = v10
    case "hello":
        var v11 int16 = 42
        _ = v11
    }
    switch -305 {
    case 9223372036854775807:
        if (!true) {
            if false {
                fmt.Println(p0)
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int8 = 127
                    p0 = 981
                    fmt.Println(fi12, v13, fi12)
                    _ = fi12
                    _ = v13
                }
            }
            if (0.1 == (0.0 * 3.14159)) {
                v14 := GenS1{F0: 2147483647, F1: 0, F2: 61}
                p0 += 403
                _ = v14
            }
            if (!(2 == 2)) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    fmt.Println(p0, p0)
                    v16 := GenS1{F0: 1, F1: 966, F2: -599}
                    var v17 int8 = -86
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 float64 = 0.0
                _ = v18
            }
        } else {
            if false {
                var v19 uint = 703
                v20 := []int{32768}
                v19 = 65535
                _ = v19
                _ = v20
            }
        }
    case 667:
        v21 := GenS1{F0: 65535, F1: 629, F2: 306}
        _ = v21
    }
    var v22 int = -992
    _ = p0
    _ = v22
    return v22
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 bool = (true && ((0.1 != -2.25) && false))
    var v23 uint8 = 44
    v1[-85] = "a"
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint8 = v23
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := []int{fi24, 32768, (fi26 | genFn0(1)), (fi26 | fi24)}
            switch "foo bar" {
            case "hello":
                var v28 float64 = ((1.5 / 1.5) / (-2.25 - 3.14159))
                _ = v28
            case "":
                switch (fi26 >> ((42) & 63)) {
                case -1:
                    fmt.Println(v23)
                case -9223372036854775808:
                    var v29 int32 = 42
                    _ = v29
                case 0:
                    fmt.Println(v23, v25, fi26)
                }
            }
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    var v30 string = "\n\t"
    switch ((v30 + v30) + "a") {
    case "a":
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS2{F0: 1, F1: "", F2: 1}
            _ = fi31
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            if (!v3) {
                v30 = v30
            }
            v34 := []int{}
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi33
            _ = v34
        }
    }
    if false {
        v1[-2147483648] = (("" + v30) + v30)
        for fi35 := 0; fi35 < 4; fi35++ {
            v36 := []int{}
            _ = fi35
            _ = v36
        }
        switch -2147483648 {
        case 42:
            if v3 {
                switch "" {
                case "hello":
                    for fi37 := 0; fi37 < 12; fi37++ {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            var v39 int64 = -1
                            switch "hello" {
                            case "a":
                                v1[32768] = v30
                            case "\n\t":
                                v40 := []int{fi38}
                                _ = v40
                            }
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                case "foo bar":
                    v23 = v23
                default:
                    switch genFn0(18446744073709551615) {
                    case 128:
                        v41 := GenS3{F0: 2}
                        _ = v41
                    case 1:
                        v30 = (v30 + ((v30 + v30) + v30))
                    }
                }
                for fi42 := 0; fi42 < 9; fi42++ {
                    v43 := []int{(fi42 / ((fi42) | 1)), (^(fi42))}
                    _ = fi42
                    _ = v43
                }
                v44 := GenS1{F0: 42, F1: -1, F2: 9223372036854775807}
                _ = v44
            }
        }
    } else {
        v45 := GenS4{F0: 1, F1: 42, F2: 477}
        _ = v45
    }
    var v46 string = (v30 + "\n\t")
    for fi47 := 0; fi47 < 1; fi47++ {
        if false {
            switch v30 {
            case "\n\t":
                v48 := []int{fi47}
                _ = v48
            case "a":
                if (82 >= 42) {
                    var v49 float64 = (3.14159 / float64(3.14159))
                    v30 = ((v46 + "a") + "foo bar")
                    var v50 float64 = v49
                    _ = v49
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        v52 := GenS3{F0: 627}
                        _ = fi51
                        _ = v52
                    }
                }
            case "foo bar":
                v3 = ((("foo bar" + v30) + (("foo bar" + v46) + "\n\t")) < v30)
            default:
                for fi53 := 0; fi53 < 9; fi53++ {
                    v46 = "hello"
                    switch v46 {
                    case "hello":
                        v46 = ((v46 + v30) + (v30 + v46))
                    case "a":
                        v3 = true
                    case "":
                        var v54 bool = true
                        _ = v54
                    default:
                        var v55 int16 = -605
                        _ = v55
                    }
                    _ = fi53
                }
            }
            var v56 float64 = ((1.5 + 1.5) + 1.5)
            _ = v56
        }
        _ = fi47
    }
    var v57 bool = (1 == 42)
    v30 = ((v46 + "a") + v46)
    fmt.Println(v30, v57)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v46:", v46)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v46
    _ = v57
}
