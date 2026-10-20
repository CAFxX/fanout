package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 int32, p2 uint) uint64 {
    var v4 uint32 = 606
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(v4, fi5, p0)
            if false {
                v7 := []int{}
                switch (fi6 & (fi6 ^ -2147483648)) {
                case -534:
                    p0 = (p0 - p0)
                default:
                    v8 := GenS1{F0: (!false)}
                    var v9 string = "foo bar"
                    v10 := GenS1{F0: (255 <= 251)}
                    _ = v10
                    _ = v8
                    _ = v9
                }
                fmt.Println(p2)
                _ = v7
            }
            if (false || ((-(p1)) == 1)) {
                p0 = 42
                if (((p2 | p2) <= (p2 % ((683) | 1))) && true) {
                    fmt.Println(p0, p0, p1)
                    p0 -= -9223372036854775808
                    v11 := []int{(fi5 - fi5), int(v4), (fi5 + fi5)}
                    _ = v11
                }
            } else {
                v12 := GenS2{F0: "", F1: 42}
                p2 += p2
                var v13 bool = true
                _ = v12
                _ = v13
            }
            _ = fi6
        }
        fmt.Println(fi5)
        for fi14 := 0; fi14 < 18; fi14++ {
            p2 = 2
            _ = fi14
        }
        _ = fi5
    }
    if true {
        if (249 >= 32767) {
            var v15 uint16 = 573
            if false {
                v16 := GenS3{F0: (v15 * (v15 | 329)), F1: (p1 & p1), F2: (^(v4))}
                _ = v16
            } else {
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 uint = 42
                    fmt.Println(v18)
                    _ = fi17
                    _ = v18
                }
            }
            var v19 int16 = 1000
            _ = v15
            _ = v19
        } else {
            p0 -= p0
            fmt.Println(p1, p1)
        }
        switch 2147483647 {
        case 4294967295:
            var v20 uint64 = 18446744073709551615
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if false {
            switch fi21 {
            case 4294967295:
                v4 -= (v4 + v4)
            case -803:
                if (!false) {
                    if true {
                        v4 = 4294967295
                    }
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v23 := []int{}
                        v24 := GenS1{F0: (p1 < 2147483647)}
                        v23 = append(v23, (fi22 | -1))
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(p2, v4, v4)
                }
            }
            fmt.Println(v4)
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 int16 = 0
                v27 := GenS2{F0: "a", F1: 42}
                var v28 int8 = 89
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 int16 = -1
    var v2 uint = 18446744073709551615
    var v3 bool = true
    fmt.Println(v1)
    var v29 int8 = 0
    v30 := []int{}
    v31 := GenS4{F0: 0}
    var v32 uint32 = 4294967295
    fmt.Println(v1)
    var v33 uint32 = v32
    var v34 string = "a"
    v1 = v1
    v34 = v34
    if v3 {
        if len(v30) > 0 {
            v30[0] = 659
        }
    } else {
        v34 = v34
        var v35 uint = v2
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = v35
    }
    var v36 uint32 = (v32 + (v32 >> ((v2) & 63)))
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    switch -798 {
    case 1:
        if (!(1e10 == (1e10 + ((0.0 - ((float64(1.5) * 1e10) / 1e10)) * 0.1)))) {
            var v37 float32 = (1e10 / float32(1.5))
            var v38 int32 = -2147483648
            v33 -= v33
            _ = v37
            _ = v38
        } else {
            switch "" {
            case "a":
                var v39 bool = (0 >= 1)
                _ = v39
            case "\n\t":
                if false {
                    var v40 float64 = 0.0
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v3 = (32767 != v1)
                        _ = fi41
                    }
                }
            }
            var v42 uint16 = 255
            _ = v42
        }
    case -2147483648:
        for fi43 := 0; fi43 < 12; fi43++ {
            fmt.Println(v2, v36)
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
}
