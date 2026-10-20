package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint8) int16 {
    var v7 string = ("foo bar" + (("" + "\n\t") + "hello"))
    v8 := []int{243}
    v7 = ("" + "foo bar")
    _ = p0
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 int64, p1 bool, p2 int64) uint64 {
    switch 1 {
    case 566:
        var v9 uint32 = 300
        _ = v9
    default:
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                p2 = p2
                p0 -= -9223372036854775808
                _ = fi10
            }
            var v11 int16 = 32767
            if (!(1 <= 18446744073709551615)) {
                p0 += 523
                if (!false) {
                    fmt.Println(p1)
                    var v12 float32 = (float32(1e10) * 0.0)
                    _ = v12
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        fmt.Println(v11, v11)
                        var v15 bool = (!(("hello" <= "foo bar") && false))
                        _ = fi14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch fi16 {
                        case 42:
                            v17 := GenS1{F0: 419}
                            _ = v17
                        case -520:
                            v18 := []int{(^(fi13)), fi16}
                            _ = v18
                        }
                        var v19 float64 = 0.1
                        var v20 float32 = 0.1
                        _ = fi16
                        _ = v19
                        _ = v20
                    }
                    switch fi13 {
                    case 4294967295:
                        v21 := GenS1{F0: 426}
                        _ = v21
                    case -609:
                        p1 = true
                    }
                    _ = fi13
                }
            } else {
                p1 = (!p1)
                if (0 > 2) {
                    switch 541 {
                    case 32767:
                        v22 := []int{-694, -260, 218}
                        _ = v22
                    case 752:
                        v23 := []int{-721}
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 int = (fi24 - (fi24 ^ (-(fi24))))
                        _ = fi24
                        _ = v25
                    }
                }
            }
            _ = v11
        }
    }
    var v26 int32 = 980
    v27 := GenS1{F0: 599}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 904
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 42
    var v3 uint8 = 1
    var v4 int16 = 0
    var v5 uint = 65535
    v6 := make(map[string]bool)
    v6[("\n\t" + ("hello" + ("a" + "foo bar")))] = (128 != (v3 - (v3 - (v3 / ((uint8(v5)) | 1)))))
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            if false {
                v30 := GenS1{F0: 1}
                v3 -= (^(v3))
                _ = v30
            } else {
                if true {
                    var v31 uint = 4294967295
                    v2 += (v2 >> ((v5) & 63))
                    var v32 uint8 = (v3 << ((v5) & 63))
                    _ = v31
                    _ = v32
                } else {
                    v1[(-(fi28))] = fi28
                    v33 := []int{}
                    switch (("a" + "hello") + ("" + "hello")) {
                    case "foo bar":
                        v2 -= v2
                    }
                    _ = v33
                }
                fmt.Println(v5, fi29, v2)
            }
            v34 := GenS1{F0: (v2 & uint64(v5))}
            _ = fi29
            _ = v34
        }
        _ = fi28
    }
    var v35 int8 = 80
    switch "" {
    case "hello":
        var v36 bool = true
        _ = v36
    case "a":
        if false {
            fmt.Println(v35, v4)
        }
    case "\n\t":
        v3 = v3
    }
    var v37 string = ""
    v38 := []int{-23, 32768, 42}
    var v39 int8 = int8(v2)
    _ = len(v1)
    var v40 uint32 = 65535
    if true {
        for fi41 := 0; fi41 < 20; fi41++ {
            if (1 >= -657) {
                var v42 uint8 = v3
                _ = v42
            }
            switch fi41 {
            case 2147483647:
                switch (v37 + ("\n\t" + v37)) {
                case "":
                    v4 -= v4
                }
            case 1:
                fmt.Println(v5, fi41, v2)
            case 42:
                fmt.Println(v5)
            default:
                var v43 int64 = 9223372036854775807
                _ = v43
            }
            _ = fi41
        }
        var v44 int32 = 1
        _ = v44
    } else {
        v45 := GenS1{F0: uint64(v5)}
        _ = v45
    }
    delete(v6, "foo bar")
    switch 128 {
    case 280:
        v37 = "hello"
    default:
        var v46 uint64 = v2
        _ = v46
    }
    fmt.Println(v3)
    var v47 float64 = (1e10 / 3.14159)
    v3 = (v3 << ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
