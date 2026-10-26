package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    var v3 uint8 = 17
    var v4 uint = 352
    var v5 float64 = 1e10
    var v6 uint = v4
    switch "foo bar" {
    case "hello":
        delete(v2, 2147483648)
    case "a":
        v4 = (v6 << ((v6) & 63))
    case "\n\t":
        var v7 float64 = v5
        _ = v7
    }
    v5 = v5
    var v8 int16 = -32768
    switch "a" {
    case "hello":
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = (true && true)
            if v10 {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 bool = true
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 string = ""
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    v4 = 2
                    _ = fi15
                }
            } else {
                v3 -= (v3 % ((0) | 1))
                if v10 {
                    v16 := []int{32767, (fi9 - (fi9 - fi9)), 579}
                    _ = v16
                } else {
                    v17 := []int{}
                    _ = v17
                }
            }
            _ = fi9
            _ = v10
        }
    case "a":
        switch 1 {
        case -421:
            v18 := GenS1{F0: 2147483648}
            _ = v18
        case 9223372036854775807:
            switch -224 {
            case 33:
                v19 := GenS1{F0: uint32(v3)}
                _ = v19
            case 32767:
                if false {
                    if true {
                        var v20 int = 1
                        _ = v20
                    } else {
                        if true {
                            fmt.Println(v1, v5)
                            v4 = (v4 << ((65535) & 63))
                        }
                        v21 := GenS1{F0: 42}
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            v22 := v2[2147483648]
                            _ = v22
                        }
                        _ = v21
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            fmt.Println(fi23, v5)
                            if true {
                                var v25 float64 = v5
                                var v26 uint32 = 65535
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int = (-(fi24))
                                _ = v27
                            }
                            _ = fi24
                        }
                        _ = fi23
                    }
                }
            case -851:
                var v28 float64 = 1.5
                _ = v28
            }
        case -2147483648:
            var v29 int32 = 1
            _ = v29
        }
    case "\n\t":
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 float32 = 1e10
            v32 := []int{fi30, (fi30 | (fi30 | (-(fi30))))}
            _ = fi30
            _ = v31
            _ = v32
        }
    default:
        if (127 == 0) {
            var v33 int64 = 9223372036854775807
            var v34 uint32 = 673
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS1{F0: 718}
        fmt.Println(v5, v4, v5)
        if true {
            if ((0 >= 65535) || ((true && ((-2147483648 <= 42) || true)) || false)) {
                var v37 float32 = ((1e10 - (float32(v5) - v1)) / v1)
                v4 = (v6 & v4)
                for fi38 := 0; fi38 < 12; fi38++ {
                    v39 := GenS2{F0: 2, F1: 1}
                    v40 := GenS1{F0: 967}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = v37
            }
        } else {
            v4 = 2
            v1 -= 0.1
        }
        _ = fi35
        _ = v36
    }
    switch 32768 {
    case 42:
        var v41 int64 = -291
        _ = v41
    case -856:
        v42 := GenS3{F0: 0, F1: true, F2: true}
        _ = v42
    default:
        v43 := GenS2{F0: (^(v6)), F1: 1}
        _ = v43
    }
    delete(v2, 651)
    v44 := GenS4{F0: 0}
    _ = len(v2)
    v8 -= v8
    var v45 string = "\n\t"
    v46 := GenS3{F0: 42, F1: ("a" >= v45), F2: false}
    var v47 uint32 = 0
    switch 127 {
    case -607:
        if false {
            v48 := GenS3{F0: 0, F1: ((v1 != (0.1 * v1)) && true), F2: false}
            switch "\n\t" {
            case "foo bar":
                v49 := GenS5{F0: v8, F1: false}
                _ = v49
            }
            _ = v48
        } else {
            var v50 uint = v6
            _ = v50
        }
    case 0:
        fmt.Println(v1, v3, v5)
    case 1:
        var v51 bool = (!true)
        _ = v51
    default:
        v45 = ((v45 + v45) + (v45 + v45))
    }
    if false {
        var v52 bool = false
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v8
}
