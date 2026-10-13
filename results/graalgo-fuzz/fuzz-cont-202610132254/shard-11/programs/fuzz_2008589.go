package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    for fi4 := 0; fi4 < 19; fi4++ {
        v2 -= (v2 / ((v2) | 1))
        _ = fi4
    }
    var v5 uint32 = 4294967295
    v3["\n\t"] = false
    var v6 int8 = 0
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float32 = (-2.25 - 0.1)
        var v9 int64 = -1
        var v10 uint = v1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int8 = (v6 & 100)
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := []int{4294967295}
        var v14 float64 = (0.0 / 0.1)
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            if false {
                var v17 uint16 = 1
                v11 += v6
                fmt.Println(fi15)
                v18 := []int{(-(fi15)), 42, 127, fi16}
                switch fi15 {
                case 32768:
                    v11 += v6
                case 4294967295:
                    var v19 int8 = (v6 >> ((v1) & 63))
                    _ = v19
                }
                _ = v17
                _ = v18
            }
            v5 = 42
            _ = fi16
        }
        v20 := GenS1{F0: 284, F1: true}
        _ = fi15
        _ = v20
    }
    if false {
        var v21 string = ("hello" + "a")
        var v22 int = 357
        for fi23 := 0; fi23 < 2; fi23++ {
            _ = len(v3)
            v2 -= v2
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        var v24 string = ("a" + ("foo bar" + ("foo bar" + "a")))
        v25 := GenS2{F0: -32768, F1: 42, F2: (v24 + "foo bar")}
        _ = v24
        _ = v25
    }
    if (true && (!false)) {
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 4; fi27++ {
                if true {
                    v11 = v11
                    fmt.Println(v1, v1, fi27)
                }
                var v28 float32 = (-2.25 * 0.1)
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                v1 = v1
                if true {
                    if false {
                        v30 := GenS1{F0: v1, F1: true}
                        fmt.Println(v2, v2, v11)
                        var v31 uint8 = 60
                        _ = v30
                        _ = v31
                    } else {
                        v6 = (v6 % ((-1) | 1))
                        if ((fi29 | (fi29 & fi26)) >= (fi29 >> ((255) & 63))) {
                            var v32 string = ((("foo bar" + "\n\t") + ("hello" + "foo bar")) + "")
                            var v33 int64 = -1
                            fmt.Println(v11, v1, fi29)
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = len(v3)
                }
                var v34 int64 = int64(fi26)
                _ = fi29
                _ = v34
            }
            _ = fi26
        }
    } else {
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS1{F0: 4294967295, F1: (0.1 <= ((1.5 * (0.0 - ((0.0 - (-2.25 + 3.14159)) * 0.0))) * float32(3.14159)))}
            _ = fi35
            _ = v36
        }
        var v37 bool = (-325 != -1)
        v11 += v6
        _ = v37
    }
    if false {
        v5 = (^(v5))
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := []int{(fi38 / (((fi38 / ((fi38) | 1))) | 1))}
            var v40 uint32 = v5
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        var v41 uint = v1
        _ = v41
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            v1 = (^(v1))
        }
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        if false {
            v1 -= v1
            if false {
                var v43 uint8 = 42
                _ = v43
            }
        } else {
            var v44 int16 = -1
            if ((-1 > -810) || true) {
                v44 = (v44 - (v44 | (v44 ^ (^(v44)))))
            } else {
                fmt.Println(v1, v1, v44)
                v11 = v11
                v11 -= -1
            }
            var v45 float64 = (0.0 / 1.5)
            _ = v44
            _ = v45
        }
        if (!(!false)) {
            v46 := GenS3{F0: v6}
            v47 := GenS2{F0: 42, F1: 1000, F2: "\n\t"}
            switch "a" {
            case "foo bar":
                var v48 int16 = 0
                _ = v48
            case "\n\t":
                v1 -= 10
            default:
                var v49 uint32 = v5
                _ = v49
            }
            _ = v46
            _ = v47
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
