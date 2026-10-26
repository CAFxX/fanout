package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int16; F2 bool }
func genFn0() int64 {
    var v6 uint64 = 2
    if (!(42 >= -2147483648)) {
        v7 := []int{-231}
        _ = len(v7)
        _ = v7
    } else {
        fmt.Println(v6)
    }
    _ = v6
    return -1
}

func main() {
    v1 := []uint32{4294967295, 65535, 883, 2147483648}
    v2 := make(map[uint32]string)
    var v3 float64 = 0.0
    var v4 int = 816
    var v5 uint64 = 0
    v5 = 0
    switch "hello" {
    case "":
        if false {
            if (v5 >= 777) {
                var v8 int8 = 111
                if (42 >= 0) {
                    fmt.Println(v3, v5, v5)
                } else {
                    v5 = v5
                    var v9 uint8 = 87
                    fmt.Println(v4)
                    _ = v9
                }
                _ = v8
            }
            var v10 int32 = -2147483648
            _ = v10
        }
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                _ = len(v2)
                v5 = (^(v5))
            }
            var v12 uint = 0
            v12 += 643
            _ = fi11
            _ = v12
        }
    default:
        if true {
            v4 = v4
        } else {
            var v13 int16 = -49
            v14 := GenS1{F0: 0}
            v15 := GenS2{F0: 65535, F1: 16}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    var v16 int = v4
    var v17 uint16 = 32767
    if (1 > 2) {
        v18 := GenS3{F0: 714}
        _ = v18
    } else {
        if false {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 float64 = float64(-2.25)
                switch (v16 - int(v5)) {
                case 2147483647:
                    v3 = ((v3 * (-2.25 * 0.1)) + v20)
                case -502:
                    v21 := v2[72]
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            v17 -= (v17 & v17)
            v22 := GenS1{F0: -9223372036854775808}
            _ = v22
        } else {
            var v23 bool = false
            var v24 int32 = 1
            _ = v23
            _ = v24
        }
        v25 := GenS3{F0: 4294967295}
        _ = v25
    }
    switch (v16 / (((-(v16))) | 1)) {
    case -333:
        for fi26 := 0; fi26 < 2; fi26++ {
            v17 += v17
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 int16 = -32768
                _ = fi27
                _ = v28
            }
            var v29 int = 128
            _ = fi26
            _ = v29
        }
    case -550:
        var v30 int32 = 42
        _ = v30
    case 2147483647:
        v2[65535] = ("hello" + "a")
    }
    var v31 uint = 408
    v32 := GenS3{F0: 42}
    v31 = 739
    v4 = -9223372036854775808
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 int8 = -115
            if true {
                var v36 string = ""
                _ = v36
            } else {
                switch ("a" + "") {
                case "a":
                    _ = len(v1)
                case "\n\t":
                    switch ("\n\t" + (("" + "\n\t") + "\n\t")) {
                    case "\n\t":
                        if (0 > uint32(v17)) {
                            v3 -= v3
                        } else {
                            v37 := GenS3{F0: 0}
                            fmt.Println(v5, fi34, v4)
                            _ = v37
                        }
                    case "hello":
                        v31 -= v31
                    case "":
                        var v38 uint32 = 241
                        v16 = fi34
                        _ = v38
                    default:
                        var v39 uint = v31
                        _ = v39
                    }
                default:
                    var v40 uint64 = (^(v5))
                    _ = v40
                }
                if false {
                    var v41 uint16 = v17
                    var v42 float32 = 0.0
                    v5 = v5
                    v31 = v31
                    v4 += v4
                    _ = v41
                    _ = v42
                } else {
                    var v43 string = "\n\t"
                    _ = v43
                }
            }
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 int16 = 32767
        if (v45 < (v45 % (((v45 << (((v31 >> (((v31 << ((v31) & 63))) & 63))) & 63))) | 1))) {
            v46 := GenS3{F0: uint32(v5)}
            _ = v46
        } else {
            v47 := GenS3{F0: 768}
            var v48 uint32 = 2147483648
            v48 += (v48 - v48)
            _ = v47
            _ = v48
        }
        v49 := GenS4{F0: 2, F1: (v45 / (((v45 - (^(v45)))) | 1)), F2: (!(false && (1 > 17)))}
        _ = fi44
        _ = v45
        _ = v49
    }
    switch "foo bar" {
    case "hello":
        for fi50 := 0; fi50 < 6; fi50++ {
            for fi51 := 0; fi51 < 2; fi51++ {
                fmt.Println(v5, fi50, v5)
                v5 -= v5
                var v52 int = fi51
                _ = fi51
                _ = v52
            }
            var v53 uint8 = 42
            v16 = fi50
            _ = fi50
            _ = v53
        }
    case "a":
        v4 -= v4
    case "foo bar":
        var v54 bool = false
        _ = v54
    }
    var v55 int = -1
    for fi56 := 0; fi56 < 8; fi56++ {
        if (0 > (^(v31))) {
            v57 := GenS4{F0: 2, F1: 1000, F2: true}
            _ = v57
        }
        v58 := GenS4{F0: 769, F1: -788, F2: (v5 <= 1)}
        v59 := GenS3{F0: 151}
        _ = fi56
        _ = v58
        _ = v59
    }
    v5 = v5
    var v60 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v55
    _ = v60
}
