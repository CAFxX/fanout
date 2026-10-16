package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0() int8 {
    v4 := []int{}
    var v5 int = 32768
    v6 := GenS1{F0: 572, F1: 554}
    _ = v4
    _ = v5
    _ = v6
    return -109
}

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 255
    var v3 string = ""
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 -= v2
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{fi7, fi7, 1}
            var v10 uint32 = 2147483648
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    switch -481 {
    case 32768:
        if (0 >= 42) {
            if (42 >= 2) {
                fmt.Println(v2, v2, v2)
                for fi11 := 0; fi11 < 19; fi11++ {
                    if ((v2 % (((^(v2))) | 1)) > 0) {
                        for fi12 := 0; fi12 < 1; fi12++ {
                            var v13 uint8 = v1
                            if true {
                                v14 := GenS2{F0: 127, F1: 2, F2: 0}
                                _ = v14
                            }
                            if (true && (!(42 <= (-(fi12))))) {
                                v15 := GenS3{F0: 2}
                                v1 += 255
                                v16 := []int{(fi12 % ((128) | 1)), (fi11 >> ((0) & 63)), 399, -9223372036854775808}
                                _ = v15
                                _ = v16
                            } else {
                                v3 = ""
                                v1 = (v1 | v1)
                                v13 += (^(v13))
                            }
                            _ = fi12
                            _ = v13
                        }
                        var v17 uint = 255
                        v1 += (v1 ^ (v1 | 187))
                        _ = v17
                    } else {
                        fmt.Println(v1)
                        var v18 uint = 2
                        _ = v18
                    }
                    _ = fi11
                }
            } else {
                v3 = v3
                var v19 float64 = 0.0
                var v20 int8 = genFn0()
                _ = v19
                _ = v20
            }
        }
    }
    v21 := []int{}
    if len(v21) > 0 {
        v21[0] = 557
    }
    fmt.Println(v3)
    var v22 int8 = 0
    v22 = 127
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            switch (fi24 >> ((131) & 63)) {
            case 2147483647:
                v25 := GenS3{F0: 42}
                _ = v25
            case -2147483648:
                var v26 int = fi23
                _ = v26
            case -9223372036854775808:
                switch (((v3 + v3) + v3) + v3) {
                case "a":
                    fmt.Println(v3)
                case "\n\t":
                    var v27 uint32 = 686
                    _ = v27
                case "":
                    v28 := GenS2{F0: v22, F1: v2, F2: genFn0()}
                    _ = v28
                }
            default:
                fmt.Println(v2)
            }
            v22 = genFn0()
            _ = fi24
        }
        fmt.Println(v3, fi23, v3)
        _ = fi23
    }
    v29 := GenS2{F0: (v22 >> ((42) & 63)), F1: 1, F2: v22}
    for fi30 := 0; fi30 < 17; fi30++ {
        fmt.Println(v1, fi30)
        if true {
            v3 = (v3 + "\n\t")
        } else {
            v31 := GenS1{F0: (fi30 * 0), F1: 1}
            v1 = v1
            _ = v31
        }
        v22 = v22
        _ = fi30
    }
    v3 = v3
    if ((true && false) && (!true)) {
        v22 = (v22 >> ((42) & 63))
    } else {
        var v32 uint16 = 1
        fmt.Println(v1)
        for fi33 := 0; fi33 < 12; fi33++ {
            switch (v3 + v3) {
            case "hello":
                var v34 uint64 = 93
                _ = v34
            case "\n\t":
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            default:
                switch v3 {
                case "foo bar":
                    switch (("hello" + "a") + "foo bar") {
                    case "a":
                        var v35 uint = 42
                        _ = v35
                    default:
                        switch fi33 {
                        case 42:
                            var v36 string = ""
                            _ = v36
                        case 4294967295:
                            v37 := GenS2{F0: 100, F1: 255, F2: genFn0()}
                            _ = v37
                        case -561:
                            for fi38 := 0; fi38 < 1; fi38++ {
                                var v39 bool = (false && false)
                                var v40 float32 = (-2.25 - -2.25)
                                fmt.Println(v40, v22)
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        }
                    }
                case "\n\t":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                default:
                    var v41 uint16 = (v2 | 478)
                    _ = v41
                }
            }
            _ = fi33
        }
        _ = v32
    }
    if len(v21) > 0 {
        v21[0] = -2147483648
    }
    if true {
        var v42 int = 127
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
}
