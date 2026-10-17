package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int8; F2 uint32 }

func main() {
    var v1 string = ""
    v2 := make(map[string]bool)
    var v3 uint16 = 42
    var v4 uint64 = 42
    var v5 int8 = -56
    var v6 int64 = 1
    var v7 int64 = v6
    v8 := GenS1{F0: true, F1: -794}
    v9 := GenS1{F0: (902 >= 85), F1: -1}
    v10 := GenS2{F0: v5}
    var v11 uint64 = v4
    var v12 uint64 = v4
    switch 128 {
    case -920:
        v13 := v2[""]
        _ = v13
    }
    v14 := GenS3{F0: 128, F1: v6}
    for fi15 := 0; fi15 < 13; fi15++ {
        if true {
            switch (v1 + "a") {
            case "a":
                fmt.Println(v4, fi15, v4)
            case "foo bar":
                for fi16 := 0; fi16 < 2; fi16++ {
                    fmt.Println(fi16, v5, v6)
                    if (((fi15 - (fi15 % ((fi16) | 1))) <= (^(fi16))) || false) {
                        v3 -= (v3 | v3)
                        v17 := GenS4{F0: v6}
                        _ = v17
                    } else {
                        v4 += (v4 & uint64(v3))
                    }
                    v18 := []int{}
                    _ = fi16
                    _ = v18
                }
            case "hello":
                v12 += (v12 >> ((65535) & 63))
            default:
                for fi19 := 0; fi19 < 2; fi19++ {
                    delete(v2, v1)
                    switch int(v3) {
                    case -364:
                        var v20 int8 = -114
                        _ = v20
                    default:
                        v21 := []int{127, -1}
                        _ = v21
                    }
                    _ = fi19
                }
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int64 = 9223372036854775807
                v6 = v23
                v12 += v12
                if false {
                    v24 := GenS1{F0: true, F1: -1}
                    if (true && false) {
                        switch (v1 + "foo bar") {
                        case "a":
                            v25 := GenS1{F0: true, F1: int16(v23)}
                            _ = v25
                        case "":
                            v26 := []int{fi22, fi22}
                            _ = v26
                        }
                        v27 := []int{-1, -1, (fi15 * -1), (^(fi22))}
                        var v28 int16 = 32767
                        _ = v27
                        _ = v28
                    } else {
                        v29 := []int{32768, (fi22 & fi22), fi22, fi15}
                        _ = v29
                    }
                    _ = v24
                } else {
                    switch 32767 {
                    case 9223372036854775807:
                        v1 = (v1 + (v1 + v1))
                    case 730:
                        if (fi22 != fi15) {
                            v5 = v5
                        } else {
                            var v30 int32 = 850
                            v11 = v12
                            _ = v30
                        }
                    }
                    if false {
                        v31 := GenS1{F0: ((-(fi22)) == (fi22 * 127)), F1: -32768}
                        var v32 int64 = v7
                        if (1.5 != (-2.25 + -2.25)) {
                            fmt.Println(v5, v23, v5)
                        } else {
                            fmt.Println(v3)
                            v2["a"] = ((v1 + v1) != (("foo bar" + "\n\t") + ""))
                        }
                        if false {
                            v1 = "hello"
                            fmt.Println(v12)
                            v33 := GenS2{F0: v5}
                            _ = v33
                        }
                        if false {
                            var v34 int8 = (^(v5))
                            _ = v34
                        }
                        fmt.Println(fi15, v4, fi15)
                        v3 = 255
                        var v35 bool = false
                        v4 = (v12 / (((v12 + (v12 * v4))) | 1))
                        _ = v31
                        _ = v32
                        _ = v35
                    }
                }
                var v36 uint64 = v4
                v37 := []int{fi22, int(v5), (fi22 << ((uint(v36)) & 63))}
                _ = fi22
                _ = v23
                _ = v36
                _ = v37
            }
        } else {
            var v38 int32 = 42
            fmt.Println(v6, v6, v3)
            _ = v38
        }
        v11 = 18446744073709551615
        for fi39 := 0; fi39 < 12; fi39++ {
            if (!true) {
                if false {
                    fmt.Println(v1, v7)
                    v6 = 42
                    delete(v2, "a")
                    if true {
                        v40 := GenS4{F0: (v6 - int64(v12))}
                        _ = v40
                    } else {
                        if ((v6 & (v6 >> ((255) & 63))) != 379) {
                            fmt.Println(v12)
                        } else {
                            if ((!(false && (!((((v1 + "") + v1) < (v1 + v1)) || true)))) || false) {
                                var v41 int8 = v5
                                _ = v41
                            }
                        }
                        var v42 float64 = 1e10
                        switch (fi15 >> ((42) & 63)) {
                        case 330:
                            fmt.Println(v4, fi39, v5)
                        case 418:
                            v12 = 433
                        case -724:
                            v43 := v2[("" + v1)]
                            _ = v43
                        }
                        _ = v42
                    }
                    var v44 int16 = -278
                    _ = v44
                }
            } else {
                v45 := GenS5{F0: v1, F1: v5, F2: 1}
                _ = v45
            }
            v6 = v6
            var v46 uint8 = 2
            _ = fi39
            _ = v46
        }
        _ = fi15
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 int32 = int32(v5)
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
