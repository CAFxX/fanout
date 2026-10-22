package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 bool = true
    var v2 int64 = -9223372036854775808
    var v3 string = "hello"
    v4 := []string{v3}
    var v5 int8 = 0
    if ((float32(3.14159) == 1e10) && (!v1)) {
        v2 = v2
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = float64(1e10)
        _ = fi6
        _ = v7
    }
    var v8 uint = 330
    if (-205 != -1) {
        if (!v1) {
            if v1 {
                var v9 int8 = v5
                _ = v9
            }
        } else {
            if v1 {
                var v10 int64 = 42
                var v11 bool = (v1 && ((!(!v1)) && true))
                for fi12 := 0; fi12 < 6; fi12++ {
                    v4 = append(v4, (v3 + ("\n\t" + ("\n\t" + "a"))))
                    v13 := GenS1{F0: 1}
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 int16 = 1000
                        _ = fi14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v10
                _ = v11
            } else {
                v3 = v3
            }
            var v16 float32 = 3.14159
            _ = v16
        }
    } else {
        v17 := GenS1{F0: uint32(v8)}
        _ = v17
    }
    if v1 {
        _ = len(v4)
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS1{F0: 2147483647}
            if v1 {
                var v20 float64 = (3.14159 + (0.1 * -2.25))
                _ = v20
            } else {
                var v21 uint32 = 879
                for fi22 := 0; fi22 < 5; fi22++ {
                    v23 := GenS2{F0: false}
                    if v1 {
                        var v24 int8 = v5
                        v25 := GenS2{F0: v1}
                        if ((fi18 & fi22) >= fi22) {
                            var v26 uint16 = 42
                            v8 = v8
                            _ = v26
                        }
                        v1 = true
                        _ = v24
                        _ = v25
                    } else {
                        v27 := GenS2{F0: v1}
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v27
                    }
                    _ = fi22
                    _ = v23
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    v1 = v1
                    var v29 bool = (!false)
                    var v30 float32 = 1e10
                    if (v29 || false) {
                        var v31 bool = ((v3 + v3) >= "hello")
                        if false {
                            var v32 uint64 = 1
                            v29 = (((96 != v5) && (255 == 2)) && true)
                            var v33 uint8 = 0
                            _ = v32
                            _ = v33
                        }
                        _ = v31
                    } else {
                        if v1 {
                            fmt.Println(v5, v30, fi28)
                        } else {
                            var v34 string = (v3 + v3)
                            v35 := GenS2{F0: ((1000 >= 1000) || (0 >= 2))}
                            _ = v34
                            _ = v35
                        }
                        var v36 int32 = -2147483648
                        fmt.Println(v21, v1, v5)
                        v3 = "hello"
                        _ = v36
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v21
            }
            for fi37 := 0; fi37 < 5; fi37++ {
                _ = len(v4)
                v5 = 1
                _ = fi37
            }
            _ = fi18
            _ = v19
        }
    } else {
        var v38 uint64 = uint64(v8)
        _ = v38
    }
    v3 = ("foo bar" + ((v3 + ("hello" + ("a" + (v3 + v3)))) + "hello"))
    var v39 uint64 = 744
    v40 := GenS3{F0: (v5 < (^(v5)))}
    fmt.Println(v2)
    var v41 int32 = -726
    v41 += -265
    var v42 string = "foo bar"
    v1 = false
    for fi43 := 0; fi43 < 19; fi43++ {
        for fi44 := 0; fi44 < 20; fi44++ {
            fmt.Println(fi43, v1, v8)
            v42 = v3
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi44
        }
        _ = fi43
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        switch v3 {
        case "hello":
            v1 = ((v2 | (^(v2))) <= (v2 ^ (v2 / (((-(v2))) | 1))))
        case "foo bar":
            if (v42 > "") {
                var v46 uint = v8
                _ = v46
            }
        case "":
            var v47 int16 = -733
            _ = v47
        }
        fmt.Println(v1, v8)
        if false {
            _ = len(v4)
            v48 := GenS3{F0: (!((fi45 + fi45) == (fi45 + fi45)))}
            if false {
                v49 := GenS4{F0: -14, F1: 1}
                v5 += (v5 ^ (-(v5)))
                _ = v49
            }
            _ = v48
        }
        _ = fi45
    }
    if len(v4) > 0 {
        v4[0] = v3
    }
    if (0 < (v39 | (v39 / (((v39 - v39)) | 1)))) {
        var v50 uint16 = 32767
        for fi51 := 0; fi51 < 4; fi51++ {
            if v1 {
                var v52 int = int(v8)
                _ = v52
            } else {
                var v53 int64 = (v2 + 9223372036854775807)
                if v1 {
                    var v54 bool = (uint8(v39) <= 1)
                    if (!(!false)) {
                        v53 += v2
                        v50 = (v50 / ((1) | 1))
                        v55 := GenS4{F0: (^(v5)), F1: -869}
                        _ = v55
                    }
                    var v56 float32 = 3.14159
                    _ = v54
                    _ = v56
                }
                _ = v53
            }
            fmt.Println(v1)
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi51
        }
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v8
}
