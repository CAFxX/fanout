package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint; F1 int8; F2 uint8 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 int8 = 0
    var v5 int8 = (v4 & v4)
    v6 := GenS1{F0: 1}
    switch "hello" {
    case "":
        if false {
            var v7 int32 = 42
            _ = v7
        } else {
            var v8 uint32 = 912
            v4 -= (v5 % ((v4) | 1))
            _ = v8
        }
    default:
        var v9 string = "\n\t"
        _ = v9
    }
    if true {
        var v10 int64 = -1
        _ = v10
    } else {
        var v11 string = ""
        _ = v11
    }
    v4 = (v4 / ((v4) | 1))
    v12 := v3["hello"]
    if true {
        var v13 uint64 = 0
        v14 := GenS2{F0: 0, F1: v13, F2: 32767}
        var v15 float32 = 0.1
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int16 = 32767
        v4 = (v4 / ((v4) | 1))
        v4 += -27
        _ = v16
    }
    v5 += 0
    var v17 int = 32768
    if (!false) {
        if (true || (-1 != int32(v17))) {
            switch ((("foo bar" + ("" + ("foo bar" + "\n\t"))) + "\n\t") + "a") {
            case "hello":
                v18 := []int{(v17 << ((715) & 63)), (^(v17))}
                _ = v18
            case "a":
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 int64 = -931
                    v12 = v12
                    var v21 bool = (float32(1e10) <= float32((1.5 + 0.1)))
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                v22 := []int{(-(v17)), -5, 42, (^(v17))}
                _ = v22
            }
            var v23 string = "hello"
            if false {
                v5 += (v4 << ((2) & 63))
            } else {
                v24 := GenS3{F0: 108, F1: false, F2: (v12 + v12)}
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint32 = 42
                    var v27 int16 = -1
                    v28 := GenS4{F0: 4294967295}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = v24
            }
            _ = v23
        }
    } else {
        switch v17 {
        case 1:
            v29 := []int{(v17 % ((v17) | 1)), (v17 * v17), (-(v17)), (v17 * v17)}
            _ = v29
        }
    }
    v17 -= v17
    fmt.Println(v4, v17, v4)
    for fi30 := 0; fi30 < 5; fi30++ {
        if (!true) {
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS2{F0: 458, F1: 822, F2: 549}
                v33 := GenS2{F0: 32767, F1: 740, F2: 59}
                var v34 uint32 = 2
                switch "" {
                case "a":
                    v17 = int(v4)
                case "hello":
                    if (float64(1e10) >= ((0.0 / -2.25) - 1e10)) {
                        var v35 int64 = 1
                        v17 = (v17 << ((255) & 63))
                        var v36 uint = 18446744073709551615
                        if false {
                            v35 -= (v35 / ((42) | 1))
                            if ((3.14159 == (1e10 + (((3.14159 + 1.5) * 0.1) * float32(float64(3.14159))))) || false) {
                                var v37 int8 = v5
                                _ = v37
                            } else {
                                v38 := GenS5{F0: 1, F1: 42}
                                _ = v38
                            }
                        }
                        _ = v35
                        _ = v36
                    }
                case "foo bar":
                    if true {
                        fmt.Println(fi31, v12)
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            var v39 uint32 = (v12 * (v34 << ((477) & 63)))
                            _ = v39
                        case "a":
                            v40 := GenS5{F0: 1000, F1: 41}
                            _ = len(v1)
                            var v41 uint32 = (v12 + (v12 / (((v34 - (^(v12)))) | 1)))
                            v42 := GenS6{F0: 4294967295, F1: (v5 * (-(v4))), F2: 127}
                            _ = v40
                            _ = v41
                            _ = v42
                        default:
                            var v43 int16 = 86
                            if (836 >= 2) {
                                v44 := GenS7{F0: 0}
                                _ = v44
                            }
                            v5 = v5
                            _ = v43
                        }
                        fmt.Println(fi30)
                    }
                default:
                    switch "hello" {
                    case "":
                        var v45 int = (fi31 ^ (fi30 / (((v17 & v17)) | 1)))
                        _ = v45
                    case "foo bar":
                        switch (-(v17)) {
                        case 128:
                            v46 := GenS1{F0: 42}
                            _ = v46
                        case 32767:
                            v17 -= 1
                        case 32768:
                            var v47 string = "hello"
                            _ = v47
                        default:
                            var v48 int16 = -1
                            _ = v48
                        }
                    case "a":
                        var v49 uint64 = 713
                        _ = v49
                    }
                    fmt.Println(v34, v17, fi31)
                }
                var v50 float32 = 1e10
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
                _ = v50
            }
            v17 = fi30
            var v51 int = (-(fi30))
            _ = v51
        } else {
            v17 -= v17
            if ((42 < 188) || false) {
                v52 := GenS4{F0: 13}
                v12 = (v12 - 1)
                _ = v52
            } else {
                var v53 bool = ((750 <= 4294967295) || (false && true))
                for fi54 := 0; fi54 < 11; fi54++ {
                    delete(v2, fi30)
                    _ = fi54
                }
                var v55 uint16 = 32767
                _ = v53
                _ = v55
            }
        }
        _ = fi30
    }
    for fi56 := 0; fi56 < 16; fi56++ {
        var v57 uint16 = 0
        if true {
            var v58 float32 = 3.14159
            var v59 float64 = -2.25
            _ = v58
            _ = v59
        }
        fmt.Println(fi56)
        _ = fi56
        _ = v57
    }
    if true {
        if true {
            fmt.Println(v5)
        }
    }
    fmt.Println(v17, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
