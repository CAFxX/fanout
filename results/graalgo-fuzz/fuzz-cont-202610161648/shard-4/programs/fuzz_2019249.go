package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint64 = 946
        _ = fi5
        _ = v6
    }
    if (false && false) {
        if (127 >= 2) {
            if (-9223372036854775808 == 1) {
                var v7 int16 = 403
                var v8 bool = (true || (!(false || true)))
                v9 := []int{-141}
                _ = v7
                _ = v8
                _ = v9
            }
            switch (("\n\t" + "\n\t") + ((((("" + "\n\t") + "foo bar") + "\n\t") + "a") + "foo bar")) {
            case "hello":
                v10 := GenS1{F0: 255}
                _ = v10
            }
        } else {
            v11 := GenS1{F0: 42}
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint8 = 255
                var v14 float64 = (0.1 + (float64(float32(float64(float32(3.14159)))) / 1.0))
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        var v15 float64 = 0.1
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 float64 = (((v15 / float64((3.14159 - float32(float64(0.0))))) * (v15 - v15)) * ((v15 - (-2.25 + v15)) * (((v15 - 0.1) + (v15 / v15)) - float64(3.14159))))
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := []int{}
                v19 = append(v19, fi16)
                var v20 uint16 = 0
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    var v21 int32 = 42
    _ = v21
    return 565
}
func genFn1(p0 uint64, p1 int32) string {
    p1 = (p1 >> ((4294967295) & 63))
    p0 = p0
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = uint16(v1)
    var v3 bool = (!(true && (true || (false && (1e10 != 1e10)))))
    v4 := []uint32{42}
    switch 871 {
    case 0:
        if (2 != 42) {
            v22 := GenS1{F0: 128}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int = fi23
                var v25 int = 63
                for fi26 := 0; fi26 < 3; fi26++ {
                    fmt.Println(v2, fi26, fi23)
                    v27 := GenS1{F0: 127}
                    switch (fi23 * v25) {
                    case 76:
                        if false {
                            if v3 {
                                v2 = uint16(v1)
                                var v28 float64 = 3.14159
                                var v29 int32 = 1
                                _ = v28
                                _ = v29
                            }
                            if (v3 && (v3 || false)) {
                                v3 = (0 == 39)
                                v30 := GenS1{F0: 0}
                                fmt.Println(v24)
                                fmt.Println(fi26, fi23, fi23)
                                _ = v30
                            } else {
                                if true {
                                    v24 = fi23
                                    v3 = false
                                } else {
                                    v2 += (v2 >> ((920) & 63))
                                    if v3 {
                                        v25 = (v25 >> ((4294967295) & 63))
                                        var v31 int16 = 0
                                        var v32 int64 = -763
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                v4 = append(v4, 4294967295)
                                switch (fi23 - fi26) {
                                case 0:
                                    if true {
                                        var v33 uint32 = 874
                                        v25 -= (v25 / (((^(v25))) | 1))
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v33
                                    } else {
                                        var v34 bool = v3
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v24 = (v25 / ((int(v2)) | 1))
                                        _ = v34
                                    }
                                case 486:
                                    v25 = v25
                                case -2147483648:
                                    v35 := GenS1{F0: 42}
                                    _ = v35
                                default:
                                    v1 = (v1 + (v1 - v1))
                                }
                            }
                        } else {
                            switch ((((("a" + "") + (("hello" + "foo bar") + "a")) + "") + "\n\t") + "foo bar") {
                            case "\n\t":
                                v36 := GenS2{F0: 65535, F1: 32767}
                                _ = v36
                            case "hello":
                                v25 = v25
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v37 bool = (814 != 4294967295)
                            v1 -= 42
                            _ = v37
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            var v38 float64 = -2.25
            _ = v38
        }
    }
    var v39 int16 = 211
    var v40 float32 = 0.0
    switch "hello" {
    case "\n\t":
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 string = "\n\t"
            fmt.Println(v39)
            v2 -= v2
            _ = fi41
            _ = v42
        }
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        v43 := GenS2{F0: genFn0(), F1: v2}
        _ = v43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        var v45 uint32 = 2147483647
        fmt.Println(v1, v3, fi44)
        _ = fi44
        _ = v45
    }
    v46 := GenS2{F0: v2, F1: 891}
    if v3 {
        switch 136 {
        case -9223372036854775808:
            fmt.Println(v3)
        }
    } else {
        fmt.Println(v3, v2, v3)
        for fi47 := 0; fi47 < 19; fi47++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 -= v1
            v48 := GenS1{F0: 36}
            _ = fi47
            _ = v48
        }
    }
    var v49 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v49
}
