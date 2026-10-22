package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 42
    v3 := make(map[int]string)
    var v4 uint = 472
    v5 := GenS1{F0: -2147483648}
    var v6 int8 = 100
    if true {
        if true {
            if true {
                for fi7 := 0; fi7 < 2; fi7++ {
                    if (127 > uint8(v2)) {
                        v6 += v6
                        var v8 uint16 = uint16(v2)
                        _ = v8
                    } else {
                        if (18446744073709551615 == v4) {
                            for fi9 := 0; fi9 < 14; fi9++ {
                                var v10 int8 = (v6 | v6)
                                v6 = (v10 ^ 1)
                                switch fi9 {
                                case 776:
                                    var v11 bool = true
                                    _ = v11
                                case 9223372036854775807:
                                    var v12 string = (("a" + "a") + "\n\t")
                                    _ = v12
                                }
                                _ = fi9
                                _ = v10
                            }
                        } else {
                            v2 -= (v2 / (((v2 ^ v2)) | 1))
                            var v13 uint8 = 0
                            fmt.Println(v13, v2, v1)
                            _ = v13
                        }
                        var v14 int64 = v1
                        switch "" {
                        case "hello":
                            var v15 int = fi7
                            _ = v15
                        default:
                            var v16 uint32 = (^(v2))
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi7
                }
                var v17 uint16 = 32767
                if true {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        v17 = v17
                        v2 += 35
                        _ = len(v3)
                        _ = fi18
                    }
                    v6 = (v6 | v6)
                    if true {
                        v19 := GenS1{F0: 1}
                        for fi20 := 0; fi20 < 6; fi20++ {
                            for fi21 := 0; fi21 < 11; fi21++ {
                                var v22 int8 = 127
                                _ = fi21
                                _ = v22
                            }
                            _ = fi20
                        }
                        _ = v19
                    }
                }
                _ = v17
            }
            switch "\n\t" {
            case "a":
                if false {
                    v6 -= (v6 ^ 67)
                    var v23 int16 = 1
                    _ = v23
                }
            case "":
                fmt.Println(v2, v4)
            case "hello":
                v6 = (v6 | (v6 ^ v6))
            }
        } else {
            v4 -= (^(v4))
            v24 := GenS1{F0: 42}
            for fi25 := 0; fi25 < 3; fi25++ {
                v26 := []int{int(v6), -9223372036854775808}
                var v27 uint16 = 347
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v24
        }
        var v28 bool = false
        _ = v28
    } else {
        v1 -= 42
        v29 := GenS2{F0: false}
        var v30 int64 = (v1 + v1)
        _ = v29
        _ = v30
    }
    var v31 uint = (^(v4))
    _ = len(v3)
    v2 = (v2 ^ v2)
    switch -99 {
    case 27:
        v32 := []int{32768}
        _ = v32
    case 127:
        v33 := GenS3{F0: -1}
        _ = v33
    }
    var v34 int64 = v1
    v35 := GenS3{F0: 0}
    v3[2147483647] = "\n\t"
    switch ("hello" + "\n\t") {
    case "hello":
        fmt.Println(v1, v31, v4)
    case "":
        var v36 string = "\n\t"
        _ = v36
    case "\n\t":
        fmt.Println(v2, v34, v1)
    default:
        _ = len(v3)
    }
    var v37 int64 = int64(v2)
    v38 := v3[556]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
