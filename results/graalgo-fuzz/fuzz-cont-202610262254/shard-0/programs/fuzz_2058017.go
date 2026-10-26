package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 string; F2 int16 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]uint32)
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 0
    v1 -= (float32(0.1) * v1)
    var v7 uint8 = v6
    var v8 uint8 = v7
    v6 -= (v8 | (v8 % (((v8 | 0)) | 1)))
    v9 := v2["\n\t"]
    v10 := GenS1{F0: 0, F1: 4294967295, F2: 32767}
    v1 -= ((3.14159 / v1) / v1)
    if (!false) {
        if true {
            var v11 float32 = (v1 * (0.1 + ((v1 * v1) / (float32(float64(v1)) / (1e10 + v1)))))
            v12 := GenS1{F0: 716, F1: 2, F2: 0}
            _ = v11
            _ = v12
        }
    }
    v6 = v8
    var v13 bool = true
    switch "" {
    case "foo bar":
        delete(v5, (v9 << ((2) & 63)))
    case "a":
        if (true && v13) {
            var v14 uint16 = 32767
            var v15 float64 = 0.1
            _ = v14
            _ = v15
        }
    default:
        var v16 int = -316
        _ = v16
    }
    var v17 int64 = (v3 * (v3 * 1))
    v8 -= uint8(v9)
    for fi18 := 0; fi18 < 3; fi18++ {
        delete(v2, "a")
        switch fi18 {
        case 32767:
            var v19 uint16 = 905
            _ = v19
        case 445:
            for fi20 := 0; fi20 < 20; fi20++ {
                if (!(v9 >= (v9 << ((499) & 63)))) {
                    var v21 float64 = 0.1
                    if false {
                        if (v17 < v3) {
                            if (fi18 != fi18) {
                                v22 := v2["foo bar"]
                                var v23 uint16 = 42
                                v24 := GenS2{F0: 195, F1: 440, F2: -78}
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            v1 -= ((v1 + v1) + v1)
                            switch 977 {
                            case 1:
                                v9 -= 42
                                var v25 int = (^(fi20))
                                if (919 > 42) {
                                    v26 := GenS3{F0: v17}
                                    fmt.Println(v9, fi20, v3)
                                    var v27 uint64 = 356
                                    fmt.Println(v13, v9, v1)
                                    _ = len(v2)
                                    var v28 uint8 = (^(v8))
                                    fmt.Println(v13)
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                var v29 int32 = 0
                                v25 = fi18
                                _ = v25
                                _ = v29
                            default:
                                fmt.Println(v1, v13)
                                v17 = (v17 / (((v3 ^ (v17 << ((42) & 63)))) | 1))
                                v30 := []int{fi18, (fi20 >> ((255) & 63))}
                                var v31 int = (fi20 - int(v8))
                                fmt.Println(v17, v31, fi18)
                                v31 += v31
                                if (v1 == float32(float64(float32(v21)))) {
                                    v1 = v1
                                    v7 -= (^(v8))
                                }
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS3{F0: (v3 + (v3 * v3))}
                            var v33 float32 = 3.14159
                            _ = v32
                            _ = v33
                        } else {
                            if v13 {
                                if v13 {
                                    var v34 string = ""
                                    if (v13 && (42 < 1)) {
                                        v35 := GenS1{F0: fi20, F1: 1, F2: 0}
                                        fmt.Println(v17)
                                        v36 := v4[(v9 * (v9 ^ 771))]
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v34 = v34
                                        v37 := GenS1{F0: (fi18 * -437), F1: 0, F2: 32767}
                                        if v13 {
                                            v38 := v5[v9]
                                            v34 = "foo bar"
                                            v39 := GenS4{F0: v3, F1: ((v34 + v34) + (v34 + (v34 + (v34 + "\n\t")))), F2: -32768}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            fmt.Println(fi18, fi18)
                                            var v40 string = ((v34 + v34) + "foo bar")
                                            v41 := []int{fi20, (fi18 + fi18), (^(fi20))}
                                            _ = v40
                                            _ = v41
                                        }
                                        v13 = v13
                                        v21 = float64(1.5)
                                        fmt.Println(v21)
                                        var v42 uint16 = 2
                                        var v43 uint16 = v42
                                        var v44 float32 = 3.14159
                                        _ = v37
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v34
                                }
                            } else {
                                var v45 uint = 65535
                                _ = v45
                            }
                        }
                        var v46 uint32 = (v9 >> ((42) & 63))
                        _ = v46
                    } else {
                        v3 += v3
                        fmt.Println(v8, v17)
                    }
                    v6 -= v7
                    _ = v21
                } else {
                    v47 := v2["hello"]
                    _ = v47
                }
                _ = fi20
            }
        case 42:
            fmt.Println(v7, v9, v7)
        }
        _ = fi18
    }
    _ = len(v4)
    v48 := GenS2{F0: 210, F1: 357, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
