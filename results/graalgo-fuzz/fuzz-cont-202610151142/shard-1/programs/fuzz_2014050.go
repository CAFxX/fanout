package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 int8) int8 {
    fmt.Println(p0)
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: true}
            p0 += 0
            _ = fi6
            _ = v7
        }
        fmt.Println(p0, p1)
    } else {
        if (false || false) {
            p1 -= (p1 << ((255) & 63))
        }
        fmt.Println(p0, p1)
        var v8 uint64 = 0
        _ = v8
    }
    var v9 int16 = 812
    v10 := []int{4294967295, -723, 32768, -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (-(p1))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    v3 := []uint32{}
    var v4 uint = 1
    var v5 string = "\n\t"
    if false {
        var v11 uint8 = 147
        switch 9223372036854775807 {
        case 2147483647:
            v2 += v2
        default:
            for fi12 := 0; fi12 < 15; fi12++ {
                switch 4294967295 {
                case 0:
                    if (false || false) {
                        v13 := GenS1{F0: true}
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v11 += v11
                            var v15 float64 = 1e10
                            var v16 float64 = 0.1
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        switch ((v5 + ("foo bar" + (("hello" + "") + v5))) + v5) {
                        case "":
                            for fi17 := 0; fi17 < 18; fi17++ {
                                var v18 bool = ((0.1 / (1.5 / 3.14159)) < float64(0.1))
                                _ = fi17
                                _ = v18
                            }
                        default:
                            v19 := GenS1{F0: (42 >= 17)}
                            _ = v19
                        }
                        _ = v13
                    }
                case -538:
                    v1["foo bar"] = 65535
                case 32767:
                    var v20 uint8 = (^(v11))
                    _ = v20
                default:
                    v5 = v5
                }
                _ = fi12
            }
        }
        _ = v11
    } else {
        v21 := GenS2{F0: 2}
        switch ("hello" + "a") {
        case "":
            var v22 string = v5
            _ = v22
        case "foo bar":
            v23 := GenS1{F0: false}
            _ = v23
        case "hello":
            var v24 uint32 = 813
            _ = v24
        default:
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int16 = 368
                _ = fi25
                _ = v26
            }
        }
        _ = v21
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint64 = 459
        switch "a" {
        case "a":
            var v29 float32 = -2.25
            _ = v29
        case "foo bar":
            switch v5 {
            case "":
                v30 := GenS2{F0: 2}
                _ = v30
            case "hello":
                fmt.Println(v2, v28, v2)
            }
        }
        _ = fi27
        _ = v28
    }
    v3 = append(v3, 954)
    var v31 int8 = -128
    v31 = v31
    v32 := GenS2{F0: 187}
    switch v5 {
    case "\n\t":
        var v33 float64 = (float64((3.14159 - 1.5)) - (1e10 - 0.0))
        _ = v33
    case "foo bar":
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 int = (fi34 << ((v4) & 63))
            _ = fi34
            _ = v35
        }
    case "hello":
        v2 = v2
    }
    v36 := GenS1{F0: (!true)}
    var v37 uint8 = 255
    fmt.Println(v2)
    _ = len(v1)
    switch v5 {
    case "a":
        v38 := GenS3{F0: (v2 - (v2 - v2))}
        _ = v38
    case "foo bar":
        if true {
            var v39 uint8 = v37
            if (!true) {
                fmt.Println(v4, v5)
            } else {
                v31 = (v31 ^ (v31 * genFn0(480, v31)))
                if false {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        switch ((v5 + v5) + "a") {
                        case "foo bar":
                            if false {
                                if (true || false) {
                                    var v41 bool = false
                                    switch fi40 {
                                    case 1:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v42 := GenS1{F0: (!v41)}
                                    _ = v41
                                    _ = v42
                                }
                                if len(v3) > 0 {
                                    v3[0] = uint32(v4)
                                }
                            }
                        case "hello":
                            if true {
                                v2 = v2
                                fmt.Println(v37, v2)
                                fmt.Println(v31)
                            }
                        }
                        _ = fi40
                    }
                    _ = len(v1)
                }
            }
            for fi43 := 0; fi43 < 8; fi43++ {
                switch (v5 + v5) {
                case "\n\t":
                    for fi44 := 0; fi44 < 2; fi44++ {
                        v39 = v39
                        _ = fi44
                    }
                case "hello":
                    fmt.Println(v39, v4, fi43)
                case "foo bar":
                    for fi45 := 0; fi45 < 7; fi45++ {
                        var v46 bool = (false && (32768 < 104))
                        _ = fi45
                        _ = v46
                    }
                default:
                    v47 := GenS1{F0: false}
                    _ = v47
                }
                _ = fi43
            }
            _ = v39
        } else {
            var v48 int32 = (v2 ^ (v2 ^ int32(v37)))
            _ = v48
        }
    case "\n\t":
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
