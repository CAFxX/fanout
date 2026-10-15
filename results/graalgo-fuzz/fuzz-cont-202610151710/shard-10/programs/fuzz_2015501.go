package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch fi4 {
        case 42:
            for fi5 := 0; fi5 < 13; fi5++ {
                v6 := GenS1{F0: -314, F1: 127, F2: 2147483648}
                if (!(!true)) {
                    if (false || (!false)) {
                        if false {
                            v7 := []int{}
                            var v8 uint32 = 2147483648
                            fmt.Println(fi4, fi4)
                            _ = v7
                            _ = v8
                        }
                        if false {
                            v9 := []int{797, (fi4 ^ 964)}
                            _ = v9
                        }
                        v10 := GenS2{F0: 42, F1: false}
                        _ = v10
                    }
                }
                _ = fi5
                _ = v6
            }
        case 128:
            if (!true) {
                for fi11 := 0; fi11 < 3; fi11++ {
                    fmt.Println(fi4, fi11)
                    var v12 int64 = 42
                    _ = fi11
                    _ = v12
                }
            } else {
                var v13 float64 = -2.25
                _ = v13
            }
        }
        var v14 int64 = -9223372036854775808
        v15 := []int{fi4, (fi4 / ((fi4) | 1))}
        _ = fi4
        _ = v14
        _ = v15
    }
    var v16 uint8 = 2
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := []int{(fi17 % ((fi17) | 1)), fi17, fi17, fi17}
        v16 -= 255
        v16 += v16
        _ = fi17
        _ = v18
    }
    var v19 uint8 = (^(v16))
    _ = v16
    _ = v19
    return 0
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]bool)
    var v3 string = "hello"
    v20 := GenS3{F0: v3, F1: (true || (!(-2.25 != 0.0))), F2: 32767}
    v1 += (v1 >> ((255) & 63))
    switch 32768 {
    case 4294967295:
        fmt.Println(v3)
    }
    v2[v1] = (("a" + "hello") > v3)
    var v21 int64 = 1
    fmt.Println(v1, v3)
    if (!false) {
        var v22 int8 = -56
        if false {
            v3 = ((v3 + ((v3 + v3) + v3)) + v3)
            v23 := v2[v1]
            for fi24 := 0; fi24 < 4; fi24++ {
                for fi25 := 0; fi25 < 5; fi25++ {
                    v26 := []int{(^(fi24)), fi25, (fi24 * (fi24 + fi25))}
                    _ = fi25
                    _ = v26
                }
                var v27 uint64 = 1
                _ = fi24
                _ = v27
            }
            _ = v23
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            switch "\n\t" {
            case "a":
                var v29 int32 = 0
                _ = v29
            default:
                v30 := GenS1{F0: 42, F1: 151, F2: 4294967295}
                _ = v30
            }
            _ = fi28
        }
        _ = v22
    }
    switch "" {
    case "":
        fmt.Println(v3, v1, v1)
    }
    fmt.Println(v3, v3)
    var v31 uint32 = 1
    fmt.Println(v31, v31)
    if (v3 < v3) {
        var v32 int32 = 1
        var v33 uint64 = 963
        _ = v32
        _ = v33
    } else {
        fmt.Println(v3, v3, v3)
        for fi34 := 0; fi34 < 4; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                v3 = (v3 + ("hello" + "foo bar"))
                if ("foo bar" <= v3) {
                    var v36 bool = (!true)
                    v1 = v31
                    var v37 uint16 = 32767
                    v31 += (v1 >> ((757) & 63))
                    if (28 == v1) {
                        v21 = (v21 % ((int64(v37)) | 1))
                        var v38 uint64 = 1
                        _ = v38
                    }
                    _ = v36
                    _ = v37
                }
                if (false || false) {
                    v1 += (v31 % (((v1 + (v1 << ((180) & 63)))) | 1))
                }
                if true {
                    var v39 float32 = -2.25
                    _ = v39
                } else {
                    var v40 uint32 = (v31 | v31)
                    var v41 int = fi34
                    _ = v40
                    _ = v41
                }
                switch (((v3 + v3) + v3) + "foo bar") {
                case "\n\t":
                    fmt.Println(fi34)
                case "hello":
                    v42 := []int{fi35}
                    _ = v42
                }
                _ = fi35
            }
            if false {
                v43 := v2[v31]
                _ = v43
            } else {
                var v44 int = (^(fi34))
                var v45 int32 = 424
                if true {
                    v21 -= (-(v21))
                } else {
                    v3 = ((((v3 + v3) + v3) + v3) + v3)
                    v44 -= v44
                    delete(v2, v31)
                }
                _ = v44
                _ = v45
            }
            switch v3 {
            case "":
                switch v3 {
                case "a":
                    var v46 bool = true
                    _ = v46
                default:
                    v47 := GenS3{F0: v3, F1: (255 <= 65535), F2: 0}
                    _ = v47
                }
            }
            _ = fi34
        }
    }
    if true {
        switch v3 {
        case "hello":
            v3 = (v3 + "\n\t")
        case "":
            v31 += v31
        case "foo bar":
            v48 := GenS3{F0: v3, F1: false, F2: -32768}
            _ = v48
        }
        v31 = v31
    } else {
        if true {
            v49 := GenS3{F0: v3, F1: (true || (!true)), F2: 32767}
            switch 4294967295 {
            case 2147483647:
                var v50 string = v3
                _ = v50
            case 42:
                var v51 float64 = (-2.25 * -2.25)
                _ = v51
            case 127:
                v52 := v2[v1]
                _ = v52
            }
            _ = v49
        }
        v53 := []int{-323}
        if (false && (1.5 != 0.0)) {
            if ((!(true && ((0 != 0) || (!(0.1 == -2.25))))) && (127 < -5)) {
                if false {
                    v21 = (v21 >> ((65535) & 63))
                } else {
                    delete(v2, 2147483647)
                    v54 := v2[(v1 & (v1 / (((v31 / ((2147483648) | 1))) | 1)))]
                    if len(v53) > 0 {
                        v53[0] = 977
                    }
                    _ = v54
                }
                for fi55 := 0; fi55 < 18; fi55++ {
                    var v56 bool = (3.14159 == 3.14159)
                    _ = fi55
                    _ = v56
                }
            }
            var v57 float64 = (0.0 * float64(1e10))
            _ = v57
        } else {
            v2[(v1 / ((v1) | 1))] = (!((0 >= 62) && (!(-9223372036854775808 != 1))))
        }
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
}
