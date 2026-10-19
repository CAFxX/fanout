package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8) string {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 string = ("hello" + "hello")
            _ = fi5
            _ = v6
        }
        if true {
            var v7 int16 = 42
            var v8 int64 = 9223372036854775807
            fmt.Println(v7, v8)
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    p0 = p0
    v9 := []int{4294967295, 2147483647}
    _ = p0
    _ = v9
    return ("\n\t" + ("a" + ("hello" + "hello")))
}
func genFn1() int8 {
    if false {
        var v10 uint64 = 0
        _ = v10
    } else {
        var v11 string = ""
        for fi12 := 0; fi12 < 7; fi12++ {
            v11 = ("\n\t" + v11)
            fmt.Println(v11)
            if (!(0 <= 32767)) {
                fmt.Println(v11, fi12)
            }
            _ = fi12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14)
                v15 := []int{fi13, 4294967295, fi13}
                switch (v11 + v11) {
                case "foo bar":
                    var v16 int = (fi14 * (fi14 >> ((65535) & 63)))
                    _ = v16
                case "a":
                    var v17 float32 = (1e10 - (1e10 * 1e10))
                    _ = v17
                case "":
                    fmt.Println(fi13)
                default:
                    var v18 uint = 671
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
            if true {
                var v19 int32 = -334
                v20 := GenS1{F0: "\n\t"}
                if (false && (true || (-1 <= 1000))) {
                    v11 = (v11 + (("a" + v11) + "foo bar"))
                }
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        _ = v11
    }
    var v21 float32 = ((-2.25 / 1.0) - (0.0 * -2.25))
    fmt.Println(v21)
    var v22 uint32 = 65535
    switch "\n\t" {
    case "":
        var v23 uint = 18446744073709551615
        _ = v23
    case "hello":
        if (!(0 < 497)) {
            if false {
                switch 0 {
                case 2147483647:
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if true {
                            v25 := []int{fi24, (fi24 - fi24)}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 14; fi26++ {
                                fmt.Println(fi24, fi26, v22)
                                var v27 int16 = -1
                                _ = fi26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 6; fi28++ {
                                var v29 uint16 = 2
                                _ = fi28
                                _ = v29
                            }
                            v30 := GenS1{F0: "a"}
                            _ = v30
                        }
                        _ = fi24
                    }
                }
            } else {
                var v31 uint32 = 2147483647
                switch ((("a" + "a") + "foo bar") + ("\n\t" + ("hello" + "a"))) {
                case "a":
                    var v32 float32 = v21
                    _ = v32
                case "":
                    var v33 int16 = 636
                    _ = v33
                }
                var v34 uint64 = 551
                _ = v31
                _ = v34
            }
            fmt.Println(v21)
        } else {
            var v35 int16 = -1
            _ = v35
        }
    case "foo bar":
        v36 := GenS1{F0: "foo bar"}
        _ = v36
    }
    _ = v21
    _ = v22
    return 1
}

func main() {
    var v1 bool = false
    var v2 int8 = -1
    var v3 uint64 = 2
    fmt.Println(v2, v1)
    var v37 uint8 = 255
    var v38 float32 = -2.25
    var v39 string = (("a" + ((("foo bar" + "a") + "foo bar") + ("foo bar" + "hello"))) + "\n\t")
    for fi40 := 0; fi40 < 19; fi40++ {
        v41 := []int{fi40}
        _ = fi40
        _ = v41
    }
    v39 = v39
    v42 := GenS2{F0: 0, F1: -503, F2: 323}
    fmt.Println(v3, v37, v3)
    v43 := GenS3{F0: 1, F1: 42, F2: 255}
    v44 := []int{1, -2147483648, -1, 32767}
    var v45 int16 = 318
    v46 := GenS4{F0: v1}
    for fi47 := 0; fi47 < 5; fi47++ {
        v45 -= (v45 * (v45 << ((65535) & 63)))
        if true {
            if (v2 == 65) {
                for fi48 := 0; fi48 < 14; fi48++ {
                    v44 = append(v44, -761)
                    _ = fi48
                }
                v45 -= v45
            } else {
                var v49 uint64 = 373
                var v50 string = ""
                _ = v49
                _ = v50
            }
            v44 = append(v44, (^(fi47)))
            v39 = "foo bar"
        } else {
            for fi51 := 0; fi51 < 11; fi51++ {
                if (int32(v45) != 1) {
                    var v52 string = (v39 + v39)
                    _ = v52
                } else {
                    if (1 >= 378) {
                        var v53 uint = 65535
                        _ = v53
                    } else {
                        v54 := GenS3{F0: (fi47 / ((-817) | 1)), F1: genFn1(), F2: 42}
                        v1 = v1
                        v39 = "\n\t"
                        _ = v54
                    }
                }
                _ = fi51
            }
            fmt.Println(v1)
            v55 := GenS1{F0: v39}
            _ = v55
        }
        _ = fi47
    }
    var v56 float32 = (float32(1e10) + (v38 + v38))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v56
}
