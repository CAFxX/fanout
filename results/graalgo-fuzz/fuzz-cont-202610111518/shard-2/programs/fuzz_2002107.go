package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
func genFn0(p0 string, p1 int32, p2 uint16) uint8 {
    v6 := GenS1{F0: 240, F1: 1, F2: 387}
    v7 := []int{4294967295, 42}
    var v8 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 102
}
func genFn1(p0 int64) bool {
    p0 -= (p0 ^ 42)
    v9 := GenS1{F0: 65535, F1: 0, F2: 2}
    if (!(!((("\n\t" + ("foo bar" + "")) <= "foo bar") && ("foo bar" > ("hello" + ("hello" + ("" + ("hello" + "a")))))))) {
        if (true && true) {
            fmt.Println(p0)
        } else {
            if false {
                fmt.Println(p0)
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: 18446744073709551615, F1: -1, F2: 4294967295}
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 uint64 = 4294967295
                    v13 = v13
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            var v14 bool = true
            _ = v14
        }
        v15 := []int{}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int16 = 1
            fmt.Println(fi16)
            switch (fi16 + 0) {
            case 42:
                p0 = int64(v17)
            case -2147483648:
                fmt.Println(p0, p0)
            case 32767:
                if false {
                    var v18 int64 = 994
                    _ = v18
                }
            }
            _ = fi16
            _ = v17
        }
    }
    var v19 string = ("" + "a")
    _ = p0
    _ = v19
    _ = v9
    return (true || false)
}

func main() {
    var v1 int = 0
    var v2 string = (("hello" + "hello") + "")
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    var v5 float64 = 1.5
    fmt.Println(v1, v5)
    v2 = v2
    v5 = 1.5
    v1 = -2147483648
    if (((127 < 0) && false) || false) {
        var v20 uint64 = 198
        v21 := GenS2{F0: 2147483647, F1: 955, F2: 56}
        v1 = (-(v1))
        _ = v20
        _ = v21
    }
    var v22 float64 = 3.14159
    fmt.Println(v2)
    for fi23 := 0; fi23 < 13; fi23++ {
        _ = len(v4)
        switch v2 {
        case "\n\t":
            var v24 int8 = -33
            _ = v24
        case "foo bar":
            var v25 string = v2
            _ = v25
        }
        fmt.Println(v2, v2, v5)
        _ = fi23
    }
    v26 := GenS3{F0: 120}
    switch v1 {
    case 179:
        var v27 float32 = 1.5
        _ = v27
    case 42:
        var v28 int32 = -2147483648
        _ = v28
    case 32767:
        if (100 == 104) {
            switch v1 {
            case 128:
                var v29 bool = (-15 <= -2147483648)
                _ = v29
            case -1:
                v30 := []int{v1, 187}
                _ = v30
            case 127:
                v31 := []int{}
                _ = v31
            }
            fmt.Println(v22)
            var v32 int16 = -32768
            _ = v32
        } else {
            fmt.Println(v1, v22, v22)
            if true {
                v33 := []int{}
                if false {
                    if false {
                        var v34 uint = 0
                        _ = v34
                    } else {
                        var v35 int16 = 915
                        var v36 string = v2
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(v1, v5)
                    if (!false) {
                        switch 969 {
                        case 127:
                            var v37 int32 = 1
                            _ = v37
                        case 871:
                            for fi38 := 0; fi38 < 13; fi38++ {
                                v39 := GenS1{F0: 537, F1: -32768, F2: 2}
                                var v40 float32 = -2.25
                                var v41 bool = true
                                _ = fi38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        case 32767:
                            for fi42 := 0; fi42 < 15; fi42++ {
                                if (!((1.5 * 0.1) <= float32(v22))) {
                                    for fi43 := 0; fi43 < 14; fi43++ {
                                        delete(v4, -289)
                                        _ = fi43
                                    }
                                    var v44 int64 = 647
                                    _ = v44
                                }
                                _ = fi42
                            }
                        }
                        fmt.Println(v5, v2, v22)
                        v2 = ""
                    }
                } else {
                    var v45 bool = ((42 == 0) && false)
                    v33 = append(v33, 704)
                    _ = v45
                }
                var v46 float64 = v5
                _ = v33
                _ = v46
            } else {
                var v47 uint16 = 32767
                v1 -= (v1 % (((v1 & -9223372036854775808)) | 1))
                v48 := v4[42]
                _ = v47
                _ = v48
            }
            for fi49 := 0; fi49 < 3; fi49++ {
                v50 := []int{v1, fi49, (v1 ^ v1), v1}
                _ = fi49
                _ = v50
            }
        }
    }
    v51 := GenS4{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
