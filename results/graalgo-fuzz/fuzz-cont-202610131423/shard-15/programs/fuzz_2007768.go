package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    var v7 float32 = (1e10 / (1.5 * ((1.5 + 1e10) + 0.0)))
    v8 := GenS1{F0: 139}
    v9 := []int{-9223372036854775808, 2147483647, -1}
    v10 := GenS1{F0: 128}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 158
}

func main() {
    var v1 int64 = -1
    var v2 float64 = 0.1
    var v3 int32 = 1
    v4 := make(map[string]bool)
    var v5 uint64 = 0
    var v6 int8 = 42
    if false {
        if (1e10 < 0.1) {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: 48}
                v13 := GenS1{F0: genFn0()}
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := []int{833, -1, 32767, -796}
            v15 := GenS1{F0: 7}
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: 96}
            _ = v16
        }
        if (61 != 128) {
            switch "hello" {
            case "hello":
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 uint16 = 859
                    for fi19 := 0; fi19 < 20; fi19++ {
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            if (false || true) {
                                v1 -= (v1 >> ((255) & 63))
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    v20 := GenS2{F0: (v6 + (v6 - v6))}
                                    v3 -= -464
                                    var v21 float64 = v2
                                    fmt.Println(v3, v21, fi19)
                                    _ = v20
                                    _ = v21
                                case "hello":
                                    var v22 string = ""
                                    _ = v22
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(fi17)
                                        default:
                                            v23 := []int{fi19, (fi17 | 127)}
                                            _ = v23
                                        }
                                    }
                                }
                            }
                        default:
                            var v24 float64 = v2
                            _ = v24
                        }
                        v25 := []int{fi17, (fi19 & (fi19 << ((4294967295) & 63))), fi19, (fi17 >> ((0) & 63))}
                        _ = fi19
                        _ = v25
                    }
                    if false {
                        if (((!(0 > (v6 & 0))) && false) && false) {
                            if ((!false) || true) {
                                switch "" {
                                case "hello":
                                    v1 += (v1 + (v1 % ((v1) | 1)))
                                }
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    v6 -= v6
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        v18 = (^(v18))
                                    case "":
                                        v27 := []int{(fi17 % ((fi26) | 1)), fi17, (fi17 ^ (fi26 << ((880) & 63)))}
                                        _ = v27
                                    }
                                    var v28 int16 = 140
                                    var v29 bool = (v18 == 0)
                                    _ = fi26
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int8 = v6
                                _ = v30
                            } else {
                                delete(v4, "\n\t")
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            default:
                _ = len(v4)
            }
            v31 := []int{128}
            _ = v31
        }
    }
    var v32 bool = false
    v3 += v3
    v32 = v32
    if (703 > (v1 & 9223372036854775807)) {
        var v33 string = "\n\t"
        if (!v32) {
            if true {
                if (32767 >= -1) {
                    if (4294967295 == 4294967295) {
                        v34 := []int{42}
                        _ = v34
                    }
                }
                var v35 int8 = v6
                _ = v35
            }
        }
        _ = v33
    }
    switch -748 {
    case 9223372036854775807:
        v36 := GenS1{F0: 2}
        _ = v36
    case -1:
        v2 -= (3.14159 / 1e10)
    case 1:
        v3 -= v3
    }
    fmt.Println(v2, v32, v6)
    var v37 uint16 = 42
    v6 = int8(v3)
    var v38 string = ((("\n\t" + "") + "foo bar") + (("\n\t" + "a") + ("hello" + "a")))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
