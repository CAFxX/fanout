package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 int; F2 uint }
type GenS5 struct { F0 string }
func genFn0(p0 bool) uint {
    if (((!(42 <= 0)) || true) || ((-9223372036854775808 != 0) || p0)) {
        if false {
            var v6 float32 = (1.5 - 0.1)
            if p0 {
                v7 := []int{286}
                v8 := GenS1{F0: 944, F1: -1}
                _ = v7
                _ = v8
            } else {
                var v9 int8 = 42
                _ = v9
            }
            switch 128 {
            case -1:
                p0 = (!p0)
            case 128:
                v10 := []int{127, -1}
                _ = v10
            }
            _ = v6
        } else {
            if p0 {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint = 829
                    fmt.Println(fi11)
                    fmt.Println(v12, fi11, v12)
                    _ = fi11
                    _ = v12
                }
                p0 = p0
                p0 = p0
            } else {
                v13 := []int{771, -9223372036854775808, -1, 32767}
                _ = v13
            }
            var v14 uint64 = 745
            _ = v14
        }
    } else {
        var v15 uint8 = 2
        switch ("hello" + "\n\t") {
        case "a":
            v16 := GenS2{F0: uint(v15)}
            _ = v16
        case "":
            if p0 {
                fmt.Println(v15, v15)
                var v17 uint8 = v15
                for fi18 := 0; fi18 < 10; fi18++ {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v20 := []int{(fi19 << ((103) & 63))}
                        _ = len(v20)
                        fmt.Println(fi18, p0)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                _ = v17
            } else {
                var v21 uint8 = v15
                v21 = 241
                _ = v21
            }
        case "hello":
            for fi22 := 0; fi22 < 12; fi22++ {
                if p0 {
                    v23 := GenS3{F0: p0, F1: 0}
                    if (781 <= 0) {
                        v15 = (v15 & 2)
                        fmt.Println(p0, v15, v15)
                    }
                    v24 := []int{}
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint = 42
                    _ = v25
                }
                var v26 uint64 = 0
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := []int{fi22, -792, fi22}
                    _ = fi27
                    _ = v28
                }
                _ = fi22
                _ = v26
            }
        }
        _ = v15
    }
    var v29 string = "\n\t"
    var v30 int32 = -1
    v30 += -1
    _ = p0
    _ = v29
    _ = v30
    return 106
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 181
    var v3 int16 = -1
    var v4 uint8 = v2
    var v5 uint64 = 422
    var v31 float32 = 3.14159
    fmt.Println(v3)
    fmt.Println(v31)
    var v32 uint64 = 1
    if (2 >= 255) {
        switch "hello" {
        case "a":
            v32 += v32
        case "hello":
            if false {
                var v33 bool = (4294967295 < -2147483648)
                for fi34 := 0; fi34 < 13; fi34++ {
                    fmt.Println(v2, v4)
                    var v35 int64 = -947
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                if false {
                    var v36 int32 = 42
                    for fi37 := 0; fi37 < 20; fi37++ {
                        var v38 uint8 = 0
                        switch -349 {
                        case 981:
                            v39 := []int{(-(fi37)), 32768, fi37, (fi37 + fi37)}
                            _ = v39
                        case -632:
                            v40 := GenS4{F0: (^(v4)), F1: fi37, F2: genFn0((!((((1e10 / 3.14159) <= (-2.25 / (1e10 / float64(((v31 + v31) * 0.0))))) || false) && (42 != -598))))}
                            _ = v40
                        }
                        var v41 uint16 = 32767
                        _ = fi37
                        _ = v38
                        _ = v41
                    }
                    delete(v1, "\n\t")
                    _ = v36
                } else {
                    v42 := []int{}
                    switch "foo bar" {
                    case "":
                        v43 := GenS3{F0: (2 < genFn0(true)), F1: 42}
                        _ = v43
                    case "a":
                        switch -9223372036854775808 {
                        case 127:
                            v44 := GenS2{F0: uint(v4)}
                            _ = v44
                        case -2147483648:
                            if len(v42) >= 2 {
                                v42 = v42[1:len(v42)]
                            }
                        case 0:
                            v45 := v1["foo bar"]
                            _ = v45
                        }
                    default:
                        v46 := GenS3{F0: true, F1: 96}
                        _ = v46
                    }
                    v4 += v4
                    _ = v42
                }
            }
        case "foo bar":
            var v47 uint16 = 323
            _ = v47
        default:
            var v48 int16 = v3
            _ = v48
        }
        v49 := GenS5{F0: ("" + "\n\t")}
        _ = v49
    }
    v50 := GenS2{F0: 42}
    var v51 uint = 1
    fmt.Println(v51)
    v52 := v1["\n\t"]
    switch ("hello" + ("hello" + "a")) {
    case "a":
        fmt.Println(v32, v3, v51)
    }
    fmt.Println(v32)
    var v53 int64 = 9223372036854775807
    _ = len(v1)
    switch "foo bar" {
    case "foo bar":
        var v54 uint8 = v4
        _ = v54
    case "\n\t":
        var v55 uint8 = v2
        _ = v55
    case "hello":
        v52 = v52
    default:
        var v56 int16 = v3
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
