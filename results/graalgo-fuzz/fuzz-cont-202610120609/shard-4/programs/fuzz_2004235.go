package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 string = (("\n\t" + "\n\t") + "foo bar")
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    v4 := make(map[int]int)
    var v5 uint = 2
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 0.1
        switch "" {
        case "foo bar":
            if (2 > 128) {
                var v9 int32 = 547
                var v10 bool = true
                v11 := GenS1{F0: 0, F1: -1, F2: -9223372036854775808}
                _ = v10
                _ = v11
                _ = v9
            } else {
                var v12 uint16 = 22
                _ = v12
            }
        case "\n\t":
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int16 = 0
                _ = fi13
                _ = v14
            }
        }
        if true {
            v15 := GenS1{F0: (^(v5)), F1: -9223372036854775808, F2: -1}
            v16 := v2[0]
            v1 = v1
            _ = v15
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    var v17 uint8 = 255
    fmt.Println(v5, v6)
    var v18 float32 = 0.1
    var v19 uint32 = 2147483648
    v6 = (v6 & 18446744073709551615)
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 uint = (v3 / ((uint(v17)) | 1))
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if ((1e10 - 0.0) != 1.5) {
            var v23 float32 = 0.1
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS2{F0: fi22, F1: fi22}
                if false {
                    var v26 uint64 = 0
                    _ = v26
                }
                v27 := GenS1{F0: (v3 >> ((401) & 63)), F1: 9223372036854775807, F2: 146}
                _ = fi24
                _ = v25
                _ = v27
            }
            fmt.Println(v17)
            v5 += v3
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 string = ("" + (((v1 + "") + (("\n\t" + v1) + (("\n\t" + (v1 + ((v1 + "") + v1))) + ""))) + ((v1 + "a") + v1)))
            _ = fi28
            _ = v29
        }
        _ = fi22
    }
    fmt.Println(v5)
    switch v1 {
    case "a":
        switch 259 {
        case -375:
            if false {
                var v30 uint = 758
                var v31 uint64 = (v6 >> ((v30) & 63))
                switch v1 {
                case "":
                    v17 = v17
                case "foo bar":
                    fmt.Println(v19)
                }
                _ = v30
                _ = v31
            }
        case 32768:
            for fi32 := 0; fi32 < 1; fi32++ {
                for fi33 := 0; fi33 < 1; fi33++ {
                    switch (-(fi32)) {
                    case 127:
                        v3 -= (v3 / ((v3) | 1))
                    case -250:
                        var v34 string = ((("" + v1) + v1) + (v1 + ""))
                        _ = v34
                    default:
                        var v35 uint = (v5 >> ((v5) & 63))
                        _ = v35
                    }
                    _ = fi33
                }
                _ = fi32
            }
        default:
            var v36 int16 = 442
            _ = v36
        }
    case "hello":
        var v37 float32 = v18
        _ = v37
    }
    v38 := []int{9223372036854775807, 32768}
    v2[(v19 | v19)] = true
    var v39 uint32 = v19
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 int = fi40
        v42 := GenS3{F0: 213}
        v6 = 451
        _ = fi40
        _ = v41
        _ = v42
    }
    if ((!false) || (0 < v3)) {
        fmt.Println(v3, v39, v19)
    } else {
        if true {
            for fi43 := 0; fi43 < 7; fi43++ {
                switch (-(fi43)) {
                case -85:
                    switch (fi43 & (fi43 % ((-199) | 1))) {
                    case 32767:
                        v6 -= v6
                    case -1:
                        var v44 uint = (v5 % ((v5) | 1))
                        _ = v44
                    case 777:
                        v45 := GenS1{F0: (v3 * v3), F1: -9223372036854775808, F2: 9223372036854775807}
                        _ = v45
                    }
                }
                switch v1 {
                case "\n\t":
                    fmt.Println(v17, v17, v17)
                case "foo bar":
                    fmt.Println(fi43)
                case "":
                    var v46 uint8 = v17
                    _ = v46
                }
                v47 := GenS1{F0: (v3 / ((1) | 1)), F1: 42, F2: -499}
                _ = fi43
                _ = v47
            }
            fmt.Println(v18)
        } else {
            v6 -= (v6 * v6)
            fmt.Println(v3, v18, v18)
        }
        var v48 int8 = -110
        _ = v48
    }
    for fi49 := 0; fi49 < 10; fi49++ {
        fmt.Println(v3)
        fmt.Println(v17)
        if false {
            if true {
                switch fi49 {
                case 2147483647:
                    _ = len(v38)
                }
                if len(v38) > 0 {
                    v38[0] = fi49
                }
                v50 := GenS1{F0: v5, F1: -1, F2: -9223372036854775808}
                _ = v50
            } else {
                fmt.Println(v3, v5, v3)
            }
            v51 := GenS1{F0: (v3 >> ((v3) & 63)), F1: 0, F2: -1}
            var v52 int32 = 873
            _ = v51
            _ = v52
        } else {
            var v53 uint32 = (v39 & 35)
            fmt.Println(v6)
            _ = v53
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
