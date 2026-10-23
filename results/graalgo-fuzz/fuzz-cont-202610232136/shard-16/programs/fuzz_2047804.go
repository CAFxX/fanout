package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0(p0 uint8, p1 uint64) int64 {
    if (!false) {
        v6 := GenS1{F0: 65535}
        _ = v6
    } else {
        if false {
            v7 := GenS1{F0: 255}
            if (p0 == (p0 >> ((447) & 63))) {
                var v8 int64 = -9223372036854775808
                var v9 int64 = 0
                _ = v8
                _ = v9
            }
            if false {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    if false {
                        var v10 float32 = -2.25
                        v11 := []int{4294967295, -804}
                        _ = v10
                        _ = v11
                    } else {
                        if false {
                            switch "\n\t" {
                            case "":
                                switch (("foo bar" + "hello") + ("\n\t" + "a")) {
                                case "\n\t":
                                    for fi12 := 0; fi12 < 18; fi12++ {
                                        var v13 int64 = 42
                                        _ = fi12
                                        _ = v13
                                    }
                                case "hello":
                                    var v14 uint8 = (p0 - (p0 + 0))
                                    _ = v14
                                default:
                                    var v15 uint8 = p0
                                    _ = v15
                                }
                            }
                        } else {
                            if true {
                                p1 = 299
                                for fi16 := 0; fi16 < 1; fi16++ {
                                    var v17 uint64 = 0
                                    p1 += 370
                                    _ = fi16
                                    _ = v17
                                }
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(p0, p0)
                                case "a":
                                    for fi18 := 0; fi18 < 14; fi18++ {
                                        p1 -= p1
                                        _ = fi18
                                    }
                                }
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    var v20 int16 = 1000
                                    v21 := GenS2{F0: 2, F1: 2, F2: 100}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            p0 -= 195
                            var v22 uint64 = p1
                            _ = v22
                        }
                        switch -2147483648 {
                        case 139:
                            for fi23 := 0; fi23 < 17; fi23++ {
                                p1 = (p1 << ((42) & 63))
                                fmt.Println(p0, p1, fi23)
                                _ = fi23
                            }
                        }
                    }
                case "hello":
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 int8 = 0
                        _ = fi24
                        _ = v25
                    }
                case "":
                    var v26 float64 = 0.1
                    _ = v26
                default:
                    fmt.Println(p0)
                }
                p1 -= (p1 & p1)
            } else {
                var v27 bool = (false || false)
                for fi28 := 0; fi28 < 13; fi28++ {
                    p1 = p1
                    _ = fi28
                }
                var v29 uint16 = 325
                _ = v27
                _ = v29
            }
            _ = v7
        }
        fmt.Println(p1, p0)
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = "a"
    v2 := []string{"foo bar"}
    var v3 float64 = 1.5
    v4 := []string{v1, "hello", v1, ("a" + ((v1 + "a") + (v1 + "")))}
    v5 := make(map[string]string)
    var v30 uint64 = 4294967295
    v3 = v3
    if (42 <= 42) {
        fmt.Println(v3, v1, v3)
    } else {
        v5[""] = v1
    }
    if false {
        var v31 int64 = 658
        _ = v31
    } else {
        var v32 string = "foo bar"
        v33 := GenS1{F0: 42}
        _ = v32
        _ = v33
    }
    var v34 int64 = 1
    var v35 int = 987
    var v36 uint64 = 1
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint16 = 42
        _ = fi37
        _ = v38
    }
    v39 := GenS3{F0: 42, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
