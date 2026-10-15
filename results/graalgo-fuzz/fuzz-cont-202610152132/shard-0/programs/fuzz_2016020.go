package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint16 = 300
        switch fi4 {
        case 4294967295:
            var v6 uint = 733
            _ = v6
        case -9223372036854775808:
            var v7 int32 = 42
            _ = v7
        case 2147483647:
            v8 := GenS1{F0: (v5 & 32767)}
            _ = v8
        default:
            var v9 int16 = 1
            _ = v9
        }
        if (!false) {
            var v10 uint64 = 18446744073709551615
            _ = v10
        } else {
            var v11 int16 = -783
            if true {
                for fi12 := 0; fi12 < 5; fi12++ {
                    if (false && true) {
                        v13 := []int{(fi4 >> ((255) & 63)), 2147483647}
                        fmt.Println(fi12, v5, v5)
                        v11 = 42
                        if false {
                            v14 := GenS1{F0: (v5 / (((^(v5))) | 1))}
                            _ = v14
                        }
                        _ = v13
                    } else {
                        v15 := []int{fi12, 32768, fi4, fi12}
                        _ = v15
                    }
                    v16 := GenS1{F0: v5}
                    _ = fi12
                    _ = v16
                }
            }
            fmt.Println(v11, fi4)
            _ = v11
        }
        _ = fi4
        _ = v5
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            if false {
                switch fi17 {
                case 0:
                    switch "foo bar" {
                    case "":
                        v19 := GenS2{F0: 189, F1: 890, F2: 1000}
                        _ = v19
                    case "a":
                        if true {
                            var v20 bool = true
                            v21 := []int{fi18}
                            v20 = (133 <= 2)
                            _ = v20
                            _ = v21
                        }
                    default:
                        v22 := GenS3{F0: 1, F1: 576}
                        _ = v22
                    }
                case 32768:
                    if false {
                        v23 := GenS2{F0: 0, F1: 84, F2: -748}
                        _ = v23
                    }
                default:
                    switch 4294967295 {
                    case 9223372036854775807:
                        var v24 int16 = 32767
                        _ = v24
                    case -2147483648:
                        var v25 uint32 = 4294967295
                        var v26 int64 = 1
                        v27 := GenS3{F0: -1, F1: 1}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 int16 = -32768
                _ = v28
            }
            fmt.Println(fi18)
            v29 := []int{(fi17 | (fi18 << ((42) & 63))), fi18}
            var v30 uint16 = 45
            _ = fi18
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int16 = 159
            v33 := GenS1{F0: 2}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi17
    }
    var v34 uint8 = 0
    v34 = (^(v34))
    _ = v34
    return -128
}

func main() {
    var v1 uint8 = 174
    var v2 bool = true
    var v3 float32 = 1e10
    if v2 {
        fmt.Println(v2, v3, v1)
        if v2 {
            var v35 uint = uint(v1)
            v1 += v1
            _ = v35
        } else {
            var v36 uint = 479
            _ = v36
        }
        if v2 {
            fmt.Println(v1)
            v37 := []int{-680}
            _ = v37
        } else {
            var v38 string = ""
            v3 = 3.14159
            v39 := GenS4{F0: v38, F1: -330}
            _ = v38
            _ = v39
        }
    }
    fmt.Println(v2, v3)
    if v2 {
        v40 := GenS5{F0: false, F1: -128, F2: 350}
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS5{F0: (525 >= 564), F1: 127, F2: 42}
            v43 := []int{fi41, fi41, fi41, -2147483648}
            v2 = v2
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = v40
    }
    v1 = v1
    if true {
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 int16 = 0
            for fi46 := 0; fi46 < 8; fi46++ {
                if (946 == 2) {
                    if (127 <= genFn0()) {
                        v45 = v45
                        var v47 int = fi44
                        _ = v47
                    } else {
                        var v48 int64 = 640
                        _ = v48
                    }
                }
                v49 := GenS3{F0: 2147483647, F1: (fi46 % ((fi44) | 1))}
                _ = fi46
                _ = v49
            }
            v50 := GenS3{F0: 2147483647, F1: fi44}
            _ = fi44
            _ = v45
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        v2 = false
        if ((v1 >> ((496) & 63)) > (v1 << ((340) & 63))) {
            v2 = (!true)
            var v52 int = (fi51 - (fi51 >> ((65535) & 63)))
            _ = v52
        } else {
            v3 -= v3
        }
        v53 := []int{fi51, (fi51 - (^(fi51)))}
        _ = fi51
        _ = v53
    }
    fmt.Println(v1)
    if v2 {
        fmt.Println(v3)
        v2 = v2
    }
    var v54 int8 = 112
    switch (("" + "a") + "") {
    case "foo bar":
        for fi55 := 0; fi55 < 11; fi55++ {
            v56 := GenS6{F0: 42, F1: 2, F2: 32767}
            var v57 int = (fi55 << ((876) & 63))
            var v58 int16 = 1000
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
    default:
        v59 := []int{-9223372036854775808, 9223372036854775807, 32768}
        _ = v59
    }
    v60 := []int{0}
    var v61 float64 = -2.25
    for fi62 := 0; fi62 < 8; fi62++ {
        var v63 string = ("\n\t" + "foo bar")
        if (0 != -2147483648) {
            var v64 uint = 18446744073709551615
            for fi65 := 0; fi65 < 14; fi65++ {
                if len(v60) > 0 {
                    v60[0] = fi62
                }
                var v66 int64 = 0
                v1 = 127
                _ = fi65
                _ = v66
            }
            _ = v64
        }
        _ = fi62
        _ = v63
    }
    v3 = 1e10
    var v67 int32 = 1
    for fi68 := 0; fi68 < 20; fi68++ {
        if v2 {
            v61 -= -2.25
        }
        for fi69 := 0; fi69 < 5; fi69++ {
            var v70 uint = 300
            v71 := GenS4{F0: "", F1: (v67 >> (((v70 ^ v70)) & 63))}
            v70 = v70
            var v72 uint = 255
            _ = fi69
            _ = v70
            _ = v71
            _ = v72
        }
        v1 += (v1 % (((^(v1))) | 1))
        _ = fi68
    }
    if len(v60) > 0 {
        v60[0] = 4294967295
    }
    v60 = append(v60, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v61:", v61)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v60
    _ = v61
    _ = v67
}
