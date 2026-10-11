package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint64 }

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -753
    v3 := make(map[uint32]string)
    v4 := []uint32{2147483647, 2, 65535, 1}
    var v5 float64 = v1
    v6 := GenS1{F0: -76}
    var v7 int32 = v2
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = 1
        if ((fi8 << ((255) & 63)) == (fi8 + (fi8 ^ (fi8 | 2147483647)))) {
            var v10 uint16 = 663
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS2{F0: 255, F1: 42, F2: -1}
                fmt.Println(v2)
                if false {
                    v2 += (v2 % (((v2 | 911)) | 1))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    _ = len(v4)
                }
                _ = fi11
                _ = v12
            }
            fmt.Println(v7, v1, v10)
            _ = v10
        } else {
            var v13 uint8 = 241
            _ = v13
        }
        var v14 float32 = 3.14159
        _ = fi8
        _ = v14
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            switch (("" + "\n\t") + "foo bar") {
            case "":
                var v16 uint32 = 65535
                _ = v16
            case "a":
                v7 += 0
            case "hello":
                var v17 uint = 4294967295
                _ = v17
            }
            v2 -= v2
            _ = fi15
        }
    }
    v2 -= 370
    v18 := GenS3{F0: 42}
    if (273 > 750) {
        v3[324] = (("a" + "a") + "a")
        switch 169 {
        case 32767:
            if true {
                var v19 uint32 = 1
                _ = v19
            }
        case 1:
            var v20 uint = 255
            _ = v20
        }
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int16 = -32768
        _ = fi21
        _ = v22
    }
    if false {
        v23 := GenS2{F0: 1, F1: -1, F2: v7}
        if true {
            v3[2147483647] = (((("" + (("" + "a") + "")) + "foo bar") + "a") + "foo bar")
            var v24 int16 = 1
            _ = v24
        }
        _ = v23
    } else {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS1{F0: -87}
            var v27 int = fi25
            v27 = (fi25 / ((-2147483648) | 1))
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = len(v4)
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                if false {
                    v7 = v7
                    fmt.Println(v7, v5, v1)
                    var v30 uint16 = 1
                    v31 := GenS3{F0: 1000}
                    switch "\n\t" {
                    case "a":
                        var v32 uint32 = 42
                        _ = v32
                    }
                    v2 -= v2
                    var v33 float32 = 3.14159
                    _ = v30
                    _ = v31
                    _ = v33
                } else {
                    if true {
                        fmt.Println(v5, v7)
                        fmt.Println(fi28, v5, v5)
                        if (303 <= 65535) {
                            v34 := GenS1{F0: 0}
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                            switch -932 {
                            case 42:
                                _ = len(v4)
                            }
                            var v35 float32 = 0.1
                            var v36 string = "\n\t"
                            _ = v34
                            _ = v35
                            _ = v36
                        } else {
                            var v37 int8 = 127
                            v5 = (v1 - v5)
                            _ = v37
                        }
                    } else {
                        var v38 uint8 = 128
                        var v39 int = (fi28 ^ (fi28 + -496))
                        fmt.Println(v2)
                        _ = v38
                        _ = v39
                    }
                }
                if false {
                    v5 = v5
                    var v40 uint8 = 0
                    _ = v40
                } else {
                    var v41 int8 = 105
                    v2 = v2
                    v42 := v3[4294967295]
                    _ = v41
                    _ = v42
                }
                _ = fi29
            }
            _ = fi28
        }
    }
    v43 := GenS1{F0: -57}
    delete(v3, 884)
    if false {
        v1 = ((v1 + (v1 + (v1 / float64(-2.25)))) - (v1 * -2.25))
        v1 = (v5 / float64(1.5))
        fmt.Println(v2, v1)
    }
    v5 = v5
    var v44 int64 = 42
    for fi45 := 0; fi45 < 19; fi45++ {
        v46 := GenS4{F0: true, F1: 18446744073709551615}
        if (0.0 <= (float64(0.0) * 0.1)) {
            v47 := GenS2{F0: 255, F1: 1, F2: v7}
            v44 = v44
            if false {
                v1 = v5
                _ = len(v4)
                var v48 int16 = -1
                _ = v48
            }
            _ = v47
        }
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
