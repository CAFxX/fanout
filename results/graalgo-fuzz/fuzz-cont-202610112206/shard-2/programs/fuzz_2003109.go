package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }

func main() {
    var v1 bool = (!((1 != 128) || true))
    v2 := make(map[uint32]string)
    v3 := []int{128}
    var v4 uint16 = 279
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: -1, F1: ("a" + "a")}
        if (!v1) {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 int = fi7
                _ = fi7
                _ = v8
            }
        } else {
            _ = len(v3)
        }
        _ = fi5
        _ = v6
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 1, F1: (("" + "hello") + ("foo bar" + "foo bar"))}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int8 = 121
        var v13 int16 = 1000
        switch (fi11 & (fi11 << ((2) & 63))) {
        case 128:
            var v14 uint32 = uint32(v4)
            _ = v14
        case -1:
            v1 = true
        default:
            var v15 bool = false
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    if false {
        var v16 uint16 = 32767
        v3 = append(v3, -9223372036854775808)
        v1 = (false || (2 > 0))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                if false {
                    var v19 int32 = 42
                    switch fi18 {
                    case 32768:
                        v19 = (v19 << ((709) & 63))
                    default:
                        fmt.Println(fi17, v1)
                    }
                    _ = v19
                }
                if (!((!(42 <= 65535)) || false)) {
                    switch "hello" {
                    case "hello":
                        switch "" {
                        case "hello":
                            v4 -= (v4 * 2)
                        }
                    case "":
                        _ = len(v2)
                    case "a":
                        v4 -= 835
                    }
                    var v20 float64 = (0.0 / 1.0)
                    v21 := GenS2{F0: 0}
                    _ = v20
                    _ = v21
                }
                _ = fi18
            }
            fmt.Println(fi17)
            v4 -= 2
            _ = fi17
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint32 = 65535
                if v1 {
                    switch fi22 {
                    case 543:
                        var v25 uint = 65535
                        _ = v25
                    case 32768:
                        switch 769 {
                        case 1:
                            v1 = (!v1)
                        default:
                            v26 := GenS1{F0: 798, F1: ("hello" + "foo bar")}
                            _ = v26
                        }
                    default:
                        var v27 int8 = 23
                        _ = v27
                    }
                    v28 := GenS1{F0: -466, F1: "\n\t"}
                    if v1 {
                        var v29 uint16 = 65535
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (!(-766 >= -1)) {
                            switch "foo bar" {
                            case "hello":
                                if (!((65535 > 699) && v1)) {
                                    var v30 int64 = -1
                                    _ = v30
                                } else {
                                    v31 := GenS1{F0: -759, F1: "foo bar"}
                                    var v32 int32 = 42
                                    v33 := GenS3{F0: v32, F1: 255, F2: 2}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            case "a":
                                var v34 int32 = -647
                                _ = v34
                            }
                            _ = len(v3)
                            v4 = 42
                        }
                        _ = v29
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v2[(^(v24))] = "hello"
                        case "a":
                            var v35 int8 = 0
                            _ = v35
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = fi23
                            }
                        }
                    }
                    v4 = 2
                    var v36 uint16 = (v4 & 2)
                    _ = v28
                    _ = v36
                }
                _ = fi23
                _ = v24
            }
        } else {
            var v37 float64 = ((1.5 + 0.1) - 3.14159)
            _ = len(v2)
            v1 = false
            _ = v37
        }
        _ = fi22
    }
    switch -1 {
    case -2147483648:
        fmt.Println(v1)
    default:
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(v1)
            v2[65535] = ((("hello" + "a") + ("foo bar" + "foo bar")) + "\n\t")
            _ = fi38
        }
    }
    var v39 uint = 1
    if ((("hello" + "\n\t") + "hello") == "foo bar") {
        if (!(!(v1 || (!(0 == 0))))) {
            var v40 uint8 = 2
            _ = v40
        } else {
            var v41 uint = v39
            v42 := GenS2{F0: -516}
            switch "" {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v41
            _ = v42
        }
    } else {
        if ((-2.25 + 1.5) == -2.25) {
            v39 = 177
            fmt.Println(v4)
            v43 := GenS1{F0: 32767, F1: "hello"}
            _ = v43
        } else {
            if (v1 && v1) {
                var v44 bool = (-612 != -112)
                fmt.Println(v44, v39, v1)
                var v45 bool = v44
                _ = v44
                _ = v45
            }
            v1 = (true || (((42 == 42) && true) || (!(32768 <= -1))))
        }
        switch 32767 {
        case 127:
            for fi46 := 0; fi46 < 10; fi46++ {
                fmt.Println(v4, v4, v1)
                for fi47 := 0; fi47 < 4; fi47++ {
                    var v48 string = ("" + ("a" + (((("a" + "\n\t") + "\n\t") + ("" + "\n\t")) + "\n\t")))
                    v1 = (127 < int8(v4))
                    var v49 uint32 = 14
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                v1 = v1
                _ = fi46
            }
        default:
            fmt.Println(v39, v4, v39)
        }
        v39 = (v39 ^ 2)
    }
    v1 = v1
    var v50 uint64 = 1
    fmt.Println(v50)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v50
}
