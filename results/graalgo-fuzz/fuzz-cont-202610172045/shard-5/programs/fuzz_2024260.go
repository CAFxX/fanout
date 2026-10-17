package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 float64 = ((float64(3.14159) / 0.1) / float64(((0.1 * 1e10) - float32((0.0 - -2.25)))))
    v4 := GenS1{F0: 127, F1: 425}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int = (^(fi5))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v3 -= float64(1.5)
        switch "" {
        case "":
            v3 = v3
        case "a":
            switch fi7 {
            case 0:
                var v8 uint64 = 2
                _ = v8
            case 42:
                if (true || false) {
                    var v9 float32 = -2.25
                    var v10 int16 = 351
                    _ = v10
                    _ = v9
                } else {
                    v3 += 1e10
                    var v11 int64 = -9223372036854775808
                    v11 -= 1
                    _ = v11
                }
            }
        }
        _ = fi7
    }
    var v12 int = 1
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS2{F0: 0}
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: -1, F1: 0}
        _ = v15
    } else {
        v16 := GenS1{F0: 1, F1: 42}
        _ = v16
    }
    var v17 int64 = -1
    v12 = v12
    for fi18 := 0; fi18 < 15; fi18++ {
        v1[fi18] = true
        v17 += v17
        _ = fi18
    }
    v3 = v3
    var v19 int8 = 0
    if (2 <= 185) {
        fmt.Println(v17, v17, v12)
        v17 = v17
        var v20 uint16 = 32767
        _ = v20
    }
    var v21 int64 = v17
    if (false || false) {
        fmt.Println(v12, v12, v12)
        for fi22 := 0; fi22 < 19; fi22++ {
            _ = len(v1)
            v23 := v2[(v12 - (fi22 - 32768))]
            if (!false) {
                v21 = v21
                v19 = v19
                for fi24 := 0; fi24 < 15; fi24++ {
                    if (2147483647 != 42) {
                        var v25 int = -959
                        if (false && false) {
                            var v26 uint64 = 2
                            _ = v26
                        } else {
                            if false {
                                var v27 int32 = -418
                                fmt.Println(v3, v23, v23)
                                _ = v27
                            }
                            var v28 int32 = -2147483648
                            switch v12 {
                            case -696:
                                v3 += v3
                            }
                            _ = v28
                        }
                        if true {
                            v29 := GenS2{F0: 2147483647}
                            _ = v29
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v30 float64 = ((3.14159 - ((3.14159 / float64(((0.1 - 3.14159) - 1.5))) * 3.14159)) - 0.0)
                                _ = v30
                            case "":
                                v31 := GenS3{F0: 425}
                                _ = v31
                            }
                        }
                        _ = v25
                    } else {
                        v19 = (v19 / ((v19) | 1))
                    }
                    v17 -= v17
                    if true {
                        var v32 int64 = -987
                        _ = v32
                    } else {
                        var v33 int16 = -471
                        _ = v33
                    }
                    _ = fi24
                }
            }
            _ = fi22
            _ = v23
        }
        v34 := GenS3{F0: 0}
        _ = v34
    } else {
        v21 += (v21 ^ v21)
        if (446 <= 496) {
            for fi35 := 0; fi35 < 10; fi35++ {
                fmt.Println(v3, v19)
                _ = fi35
            }
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS4{F0: 0, F1: -1}
                switch fi36 {
                case -769:
                    var v38 int16 = -260
                    _ = v38
                case -2147483648:
                    v12 = fi36
                }
                _ = fi36
                _ = v37
            }
        } else {
            v21 += 42
            var v39 string = ((("hello" + ("foo bar" + ((((("hello" + "a") + ("foo bar" + "foo bar")) + ("" + "")) + ("" + "")) + ""))) + "hello") + ("foo bar" + "\n\t"))
            var v40 int16 = -911
            _ = v39
            _ = v40
        }
        v17 -= (v17 >> ((1) & 63))
    }
    v41 := GenS4{F0: 809, F1: v17}
    var v42 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
