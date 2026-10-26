package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -2147483648
    var v3 string = "foo bar"
    v4 := []string{v3, ("\n\t" + v3), v3}
    v5 := make(map[string]string)
    fmt.Println(v2)
    if false {
        v2 = (v2 * 878)
        v2 += v2
    } else {
        v2 -= (v2 << ((42) & 63))
    }
    var v6 float32 = 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int = (v2 & fi7)
            var v9 uint32 = 249
            _ = fi7
            _ = v8
            _ = v9
        }
        if (2 != 1) {
            if (-2.25 >= 0.0) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch v3 {
                    case "hello":
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v6 = (-2.25 - (-2.25 / v6))
                            _ = fi11
                        }
                    }
                    v6 = v6
                    _ = fi10
                }
            } else {
                for fi12 := 0; fi12 < 9; fi12++ {
                    if true {
                        v13 := GenS1{F0: 42, F1: 0, F2: 1000}
                        _ = v13
                    }
                    _ = fi12
                }
                fmt.Println(v6, v6, v2)
            }
            var v14 string = v3
            v14 = "foo bar"
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                v2 = 32767
                if (("\n\t" + v3) != v3) {
                    switch v3 {
                    case "a":
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v6 -= float32((3.14159 * 1.5))
                            var v17 string = v3
                            v4 = append(v4, "foo bar")
                            if len(v4) > 0 {
                                v4[0] = ((v17 + "hello") + "hello")
                            }
                            _ = fi16
                            _ = v17
                        }
                    case "foo bar":
                        if true {
                            v18 := GenS1{F0: -1, F1: 42, F2: 32767}
                            v19 := GenS2{F0: 128}
                            _ = v18
                            _ = v19
                        } else {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v21 := GenS1{F0: 42, F1: -134, F2: -738}
                                    v3 = ("foo bar" + v3)
                                    v2 = fi20
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            if len(v4) > 0 {
                                v4[0] = (v3 + "\n\t")
                            }
                            var v22 uint16 = 644
                            _ = v22
                        }
                    case "":
                        v6 -= ((v6 * v6) * v6)
                    }
                    v3 = (v3 + v3)
                    switch (v2 | -541) {
                    case -789:
                        v1[65535] = -40
                    case 42:
                        v6 += (float32(1.5) + ((3.14159 * 0.0) / (v6 / float32((((1.5 + 1e10) + ((((((1.5 + -2.25) - float64(float32(1.5))) + -2.25) * ((1e10 / 1e10) - 3.14159)) - 1e10) / float64(v6))) * 3.14159)))))
                    }
                } else {
                    var v23 uint64 = 2
                    _ = v23
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    switch "" {
                    case "\n\t":
                        var v25 bool = (!false)
                        _ = v25
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                    case "":
                        v26 := GenS3{F0: false, F1: -9223372036854775808}
                        _ = v26
                    }
                    _ = fi24
                }
                _ = fi15
            }
        }
    } else {
        var v27 int = v2
        fmt.Println(v2)
        _ = v27
    }
    if false {
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float64 = float64(float32((3.14159 + 3.14159)))
            for fi30 := 0; fi30 < 1; fi30++ {
                v2 += fi28
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 uint16 = 0
        var v33 float32 = -2.25
        fmt.Println(v32)
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint32 = 371
    v35 := GenS3{F0: (128 > 0), F1: v2}
    var v36 float32 = 1e10
    v5[v3] = "\n\t"
    fmt.Println(v34, v34, v2)
    var v37 uint64 = 1
    v34 = (v34 / (((v34 << ((255) & 63))) | 1))
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 float32 = (v6 - -2.25)
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: -1}
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
