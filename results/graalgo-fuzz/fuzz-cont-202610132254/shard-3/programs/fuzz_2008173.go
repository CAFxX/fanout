package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 bool }
func genFn0(p0 int64) int {
    if true {
        v7 := []int{4294967295, 4294967295}
        if true {
            var v8 int32 = -2147483648
            v9 := GenS1{F0: "hello", F1: 255}
            _ = v8
            _ = v9
        } else {
            if true {
                p0 += 42
                switch ("hello" + "\n\t") {
                case "\n\t":
                    for fi10 := 0; fi10 < 5; fi10++ {
                        if true {
                            v11 := GenS1{F0: "hello", F1: 0}
                            _ = v11
                        }
                        _ = fi10
                    }
                case "foo bar":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v14 := GenS2{F0: 32767, F1: 0, F2: 852}
                            p0 += (p0 & (p0 + p0))
                            _ = fi13
                            _ = v14
                        }
                        v15 := GenS3{F0: 0, F1: 32767}
                        _ = fi12
                        _ = v15
                    }
                case "hello":
                    for fi16 := 0; fi16 < 9; fi16++ {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            p0 -= -52
                            fmt.Println(p0, p0)
                            _ = fi17
                        }
                        for fi18 := 0; fi18 < 11; fi18++ {
                            if ((float32(float64((3.14159 + -2.25))) - 3.14159) > 3.14159) {
                                fmt.Println(fi18, fi18)
                            } else {
                                var v19 uint32 = 65535
                                if ((!true) && (1 == v19)) {
                                    if false {
                                        switch (fi18 / ((fi16) | 1)) {
                                        case 128:
                                            switch (fi16 << ((65535) & 63)) {
                                            case 0:
                                                v20 := GenS4{F0: 255}
                                                v21 := GenS5{F0: false, F1: 0}
                                                p0 = 9223372036854775807
                                                _ = v20
                                                _ = v21
                                            case 4294967295:
                                                if (-1 != -32768) {
                                                    if (-1 != 42) {
                                                        v22 := GenS2{F0: 1, F1: 0, F2: 0}
                                                        _ = v22
                                                    }
                                                    var v23 uint16 = 32767
                                                    _ = v23
                                                }
                                            }
                                        case 2147483647:
                                            _ = len(v7)
                                        }
                                    } else {
                                        v19 += 392
                                        if false {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        }
                                        if false {
                                            if true {
                                                p0 = (p0 * -544)
                                                p0 = p0
                                            } else {
                                                v24 := GenS1{F0: "", F1: 2}
                                                if (255 <= 42) {
                                                    p0 = -9223372036854775808
                                                    var v25 uint16 = 1
                                                    if len(v7) > 0 {
                                                        v7[0] = 737
                                                    }
                                                    _ = v25
                                                }
                                                fmt.Println(fi18)
                                                var v26 uint = 1
                                                _ = v24
                                                _ = v26
                                            }
                                        }
                                        p0 = (p0 | 42)
                                    }
                                }
                                var v27 uint32 = v19
                                _ = v19
                                _ = v27
                            }
                            _ = len(v7)
                            if (75 <= 128) {
                                var v28 int = (fi16 + (^(fi18)))
                                fmt.Println(v28)
                                _ = v28
                            }
                            _ = fi18
                        }
                        _ = fi16
                    }
                }
            } else {
                v29 := GenS4{F0: 507}
                fmt.Println(p0)
                fmt.Println(p0)
                _ = v29
            }
        }
        v30 := GenS6{F0: 32767}
        _ = v30
        _ = v7
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        switch "a" {
        case "foo bar":
            v32 := GenS5{F0: false, F1: 4294967295}
            _ = v32
        }
        _ = fi31
    }
    var v33 string = ("foo bar" + "\n\t")
    _ = p0
    _ = v33
    return 9223372036854775807
}

func main() {
    var v1 int = -462
    var v2 uint8 = 64
    var v3 int8 = 1
    v4 := make(map[string]string)
    var v5 uint = 42
    v6 := []string{("\n\t" + ""), ""}
    v2 = v2
    v34 := GenS7{F0: "hello"}
    v1 += 4294967295
    v35 := GenS6{F0: -1}
    v2 = (v2 * (^(v2)))
    v36 := GenS5{F0: ((!true) && false), F1: 474}
    v37 := GenS8{F0: 32767}
    var v38 uint32 = 380
    fmt.Println(v5, v2)
    var v39 bool = (!true)
    v40 := GenS9{F0: (v5 % (((v5 % ((249) | 1))) | 1)), F1: ((v3 % (((^(v3))) | 1)) < (v3 ^ v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
